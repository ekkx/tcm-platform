package tcmsyncer_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/jobs/tcmsyncer"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock tcmrsv client
type mockTCMRSVClient struct {
	loginCalls       []tcmrsv.LoginParams
	reserveCalls     []tcmrsv.ReserveParams
	loginError       error
	reserveError     error
	myReservations   []tcmrsv.Reservation
	myReservationsError error
	getMyReservationsCallCount int
}

func (m *mockTCMRSVClient) Login(params *tcmrsv.LoginParams) error {
	m.loginCalls = append(m.loginCalls, *params)
	return m.loginError
}

func (m *mockTCMRSVClient) Reserve(params *tcmrsv.ReserveParams) error {
	m.reserveCalls = append(m.reserveCalls, *params)
	return m.reserveError
}

func (m *mockTCMRSVClient) GetMyReservations() ([]tcmrsv.Reservation, error) {
	m.getMyReservationsCallCount++
	return m.myReservations, m.myReservationsError
}

// Mock database
type mockDB struct {
	database.Execer
}

// Mock repositories
type mockReservationRepo struct {
	reservations []entity.Reservation
	getError     error
	updateCalls  []repository.UpdateReservationByIDArgs
	updateError  error
}

func (m *mockReservationRepo) GetReservationsByDate(ctx context.Context, args *repository.GetReservationsByDate) ([]entity.Reservation, error) {
	if m.getError != nil {
		return nil, m.getError
	}
	// Filter reservations by date
	var filtered []entity.Reservation
	for _, r := range m.reservations {
		if r.Date.Truncate(24*time.Hour).Equal(args.Date.Truncate(24*time.Hour)) {
			filtered = append(filtered, r)
		}
	}
	return filtered, nil
}

func (m *mockReservationRepo) UpdateReservationByID(ctx context.Context, args *repository.UpdateReservationByIDArgs) error {
	m.updateCalls = append(m.updateCalls, *args)
	return m.updateError
}

type mockUserRepo struct {
	users    map[string]*entity.User
	getError error
}

func (m *mockUserRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	if m.getError != nil {
		return nil, m.getError
	}
	user, exists := m.users[id]
	if !exists {
		return nil, pgx.ErrNoRows
	}
	return user, nil
}

func TestSyncReservationsJob_Execute(t *testing.T) {
	// Test AES key
	aesKey := []byte("12345678901234567890123456789012") // 32 bytes for AES-256

	// Encrypt test password
	encryptedPassword, err := cryptohelper.EncryptAES("testpassword", aesKey)
	require.NoError(t, err)

	// Test data
	testUser := &entity.User{
		ID:                "user1",
		EncryptedPassword: encryptedPassword,
		CreatedAt:         time.Now(),
	}

	today := time.Now().Truncate(24 * time.Hour)
	
	testReservations := []entity.Reservation{
		{
			ID:         1,
			UserID:     "user1",
			CampusType: enum.CampusTypeNakameguro,
			RoomID:     "P 427（G）",
			Date:       today.Add(10 * time.Hour), // 10:00
			FromHour:   10,
			FromMinute: 0,
			ToHour:     12,
			ToMinute:   0,
			BookerName: stringPtr("Test Booking"),
		},
		{
			ID:         2,
			ExternalID: stringPtr("already-synced"),
			UserID:     "user1",
			CampusType: enum.CampusTypeIkebukuro,
			RoomID:     "P 101",
			Date:       today.Add(14 * time.Hour), // 14:00
			FromHour:   14,
			FromMinute: 0,
			ToHour:     16,
			ToMinute:   0,
		},
		{
			ID:         3,
			UserID:     "user2",
			CampusType: enum.CampusTypeNakameguro,
			RoomID:     "P 202",
			Date:       today.Add(16 * time.Hour), // 16:00
			FromHour:   16,
			FromMinute: 30,
			ToHour:     18,
			ToMinute:   0,
		},
	}

	t.Run("successful sync", func(t *testing.T) {
		// Setup mocks
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: testReservations,
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{
				"user1": testUser,
				"user2": {
					ID:                "user2",
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
				},
			},
		}
		
		// Create job with mocked dependencies
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		// Mock tcmrsv client
		mockClient := &mockTCMRSVClient{
			myReservations: []tcmrsv.Reservation{
				{
					ID:         "ext-001",
					Campus:     tcmrsv.CampusNakameguro,
					CampusName: "中目黒・代官山キャンパス",
					Date:       today.Format("2006年01月02日"),
					RoomName:   "P 427（G）",
					TimeRange:  "10:00-12:00",
				},
				{
					ID:         "ext-002",
					Campus:     tcmrsv.CampusNakameguro,
					CampusName: "中目黒・代官山キャンパス",
					Date:       today.Format("2006年01月02日"),
					RoomName:   "P 202",
					TimeRange:  "16:30-18:00",
				},
			},
		}

		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Verify correct updates were made
		assert.Len(t, mockReservRepo.updateCalls, 2)
		
		// Check first update
		assert.Equal(t, 1, mockReservRepo.updateCalls[0].ReservationID)
		assert.Equal(t, "ext-001", *mockReservRepo.updateCalls[0].ExternalID)
		
		// Check second update
		assert.Equal(t, 3, mockReservRepo.updateCalls[1].ReservationID)
		assert.Equal(t, "ext-002", *mockReservRepo.updateCalls[1].ExternalID)

		// Verify login calls
		assert.Len(t, mockClient.loginCalls, 2)
		assert.Equal(t, "user1", mockClient.loginCalls[0].UserID)
		assert.Equal(t, "user2", mockClient.loginCalls[1].UserID)

		// Verify reserve calls
		assert.Len(t, mockClient.reserveCalls, 2)
		
		// Should call GetMyReservations once per user (2 users)
		assert.Equal(t, 2, mockClient.getMyReservationsCallCount)
	})

	t.Run("skip already synced reservations", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[1]}, // Only synced reservation
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})

		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should not make any calls
		assert.Len(t, mockClient.loginCalls, 0)
		assert.Len(t, mockClient.reserveCalls, 0)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("handle login failure", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			loginError: errors.New("invalid credentials"),
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err) // Should not fail the entire job

		// Should try to login but not update
		assert.Len(t, mockClient.loginCalls, 1)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("handle missing external ID in GetMyReservations", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			myReservations: []tcmrsv.Reservation{}, // Empty
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should reserve but not update
		assert.Len(t, mockClient.reserveCalls, 1)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("verify correct date filtering", func(t *testing.T) {
		// Include reservation from tomorrow
		allReservations := append(testReservations, entity.Reservation{
			ID:         4,
			UserID:     "user1",
			CampusType: enum.CampusTypeNakameguro,
			RoomID:     "P 303",
			Date:       today.Add(25 * time.Hour), // Tomorrow
			FromHour:   10,
			FromMinute: 0,
			ToHour:     12,
			ToMinute:   0,
		})

		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: allReservations,
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{
				"user1": testUser,
				"user2": {
					ID:                "user2",
					EncryptedPassword: encryptedPassword,
					CreatedAt:         time.Now(),
				},
			},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			myReservations: []tcmrsv.Reservation{},
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should only process today's reservations (excluding already synced)
		assert.Len(t, mockClient.reserveCalls, 2) // Only reservations 1 and 3

		// Verify all reserved dates are today
		for _, call := range mockClient.reserveCalls {
			assert.True(t, call.Date.Truncate(24*time.Hour).Equal(today))
		}
	})

	t.Run("handle decrypt password failure", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		// Use wrong key
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, []byte("wrongkey12345678"))

		mockClient := &mockTCMRSVClient{}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})

		err := job.Execute(context.Background())
		assert.NoError(t, err) // Should not fail the entire job

		// Should not make any calls
		assert.Len(t, mockClient.loginCalls, 0)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("verify same user login for multiple reservations", func(t *testing.T) {
		// Multiple reservations for same user
		sameUserReservations := []entity.Reservation{
			testReservations[0],
			{
				ID:         5,
				UserID:     "user1",
				CampusType: enum.CampusTypeNakameguro,
				RoomID:     "P 500",
				Date:       today.Add(15 * time.Hour),
				FromHour:   15,
				FromMinute: 0,
				ToHour:     17,
				ToMinute:   0,
			},
		}

		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: sameUserReservations,
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			myReservations: []tcmrsv.Reservation{
				{
					ID:         "ext-001",
					Campus:     tcmrsv.CampusNakameguro,
					CampusName: "中目黒・代官山キャンパス",
					Date:       today.Format("2006年01月02日"),
					RoomName:   "P 427（G）",
					TimeRange:  "10:00-12:00",
				},
				{
					ID:         "ext-005",
					Campus:     tcmrsv.CampusNakameguro,
					CampusName: "中目黒・代官山キャンパス",
					Date:       today.Format("2006年01月02日"),
					RoomName:   "P 500",
					TimeRange:  "15:00-17:00",
				},
			},
		}
		
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should login only once per user
		assert.Len(t, mockClient.loginCalls, 1)
		assert.Equal(t, "user1", mockClient.loginCalls[0].UserID)
		
		// Should reserve both
		assert.Len(t, mockClient.reserveCalls, 2)
		
		// Should call GetMyReservations only once after all reservations
		assert.Equal(t, 1, mockClient.getMyReservationsCallCount)
		
		// Should update both external IDs
		assert.Len(t, mockReservRepo.updateCalls, 2)
	})

	t.Run("handle reserve failure", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			reserveError: errors.New("reservation failed"),
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err) // Should not fail the entire job

		// Should try to reserve but not update
		assert.Len(t, mockClient.reserveCalls, 1)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("handle user not found", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{}, // No users
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err) // Should not fail the entire job

		// Should not make any calls
		assert.Len(t, mockClient.loginCalls, 0)
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("verify campus matching", func(t *testing.T) {
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		// Return reservation with different campus
		mockClient := &mockTCMRSVClient{
			myReservations: []tcmrsv.Reservation{
				{
					ID:         "ext-001",
					Campus:     tcmrsv.CampusIkebukuro, // Wrong campus
					CampusName: "池袋キャンパス",
					Date:       today.Format("2006年01月02日"),
					RoomName:   "P 427（G）",
					TimeRange:  "10:00-12:00",
				},
			},
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should not find matching reservation
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})

	t.Run("GetMyReservations called once even with GetMyReservations error", func(t *testing.T) {
		// Test that GetMyReservations is called only once even if it returns an error
		mockDB := &mockDB{}
		mockReservRepo := &mockReservationRepo{
			reservations: []entity.Reservation{testReservations[0]},
		}
		mockUserRepo := &mockUserRepo{
			users: map[string]*entity.User{"user1": testUser},
		}
		
		job := tcmsyncer.NewSyncReservationsJobWithDeps(mockDB, mockReservRepo, mockUserRepo, aesKey)

		mockClient := &mockTCMRSVClient{
			myReservationsError: errors.New("network error"),
		}
		job.SetTCMClientFactory(func() tcmsyncer.TCMRSVClient {
			return mockClient
		})
		
		err := job.Execute(context.Background())
		assert.NoError(t, err)

		// Should reserve successfully
		assert.Len(t, mockClient.reserveCalls, 1)
		
		// Should call GetMyReservations once
		assert.Equal(t, 1, mockClient.getMyReservationsCallCount)
		
		// Should not update any external IDs due to GetMyReservations error
		assert.Len(t, mockReservRepo.updateCalls, 0)
	})
}

func TestSyncReservationsJob_RunAt12PM(t *testing.T) {
	t.Run("verify next run time calculation", func(t *testing.T) {
		// Test at different times of day
		testCases := []struct {
			name         string
			currentTime  time.Time
			expectedWait time.Duration
		}{
			{
				name:         "morning before 12PM",
				currentTime:  time.Date(2025, 6, 10, 9, 0, 0, 0, time.FixedZone("JST", 9*60*60)),
				expectedWait: 3 * time.Hour,
			},
			{
				name:         "afternoon after 12PM",
				currentTime:  time.Date(2025, 6, 10, 15, 0, 0, 0, time.FixedZone("JST", 9*60*60)),
				expectedWait: 21 * time.Hour,
			},
			{
				name:         "exactly at 12PM",
				currentTime:  time.Date(2025, 6, 10, 12, 0, 0, 0, time.FixedZone("JST", 9*60*60)),
				expectedWait: 24 * time.Hour,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Calculate next 12PM from test time
				next12PM := time.Date(tc.currentTime.Year(), tc.currentTime.Month(), tc.currentTime.Day(), 
					12, 0, 0, 0, tc.currentTime.Location())
				if tc.currentTime.After(next12PM) || tc.currentTime.Equal(next12PM) {
					next12PM = next12PM.Add(24 * time.Hour)
				}
				
				actualWait := next12PM.Sub(tc.currentTime)
				assert.Equal(t, tc.expectedWait, actualWait)
			})
		}
	})
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}