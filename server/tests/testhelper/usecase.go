package testhelper

import (
	"context"
	"testing"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	rsv_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	mock_tcmrsv "github.com/ekkx/tcmrsv-web/server/tests/mocks/tcmrsv"
	"github.com/stretchr/testify/require"
)

const (
	TestUserID       = "testuser"
	TestUserPassword = "testpass"
	TestUserID2      = "testuser2"
	TestUserPassword2 = "testpass2"
	TestSystemID     = "testsystem"
)

// CreateTestUser creates a test user with the given ID and password
func CreateTestUser(ctx context.Context, t *testing.T, userRepo *user_repo.Repository, userID, password string) {
	encryptedPassword, err := cryptohelper.EncryptAES(password, []byte(ctxhelper.GetConfig(ctx).PasswordAESKey))
	require.NoError(t, err)

	_, err = userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
		ID:                userID,
		EncryptedPassword: encryptedPassword,
	})
	require.NoError(t, err)
}

// CreateDefaultTestUser creates a test user with default test credentials
func CreateDefaultTestUser(ctx context.Context, t *testing.T, userRepo *user_repo.Repository) {
	CreateTestUser(ctx, t, userRepo, TestUserID, TestUserPassword)
}

// CreateDefaultTestUsers creates two test users with default test credentials
func CreateDefaultTestUsers(ctx context.Context, t *testing.T, userRepo *user_repo.Repository) {
	CreateTestUser(ctx, t, userRepo, TestUserID, TestUserPassword)
	CreateTestUser(ctx, t, userRepo, TestUserID2, TestUserPassword2)
}

// GetMockTCMClient returns a mock TCM client with default login behavior
func GetMockTCMClient() *mock_tcmrsv.MockTCMClient {
	return &mock_tcmrsv.MockTCMClient{
		LoginFunc: func(params *tcmrsv.LoginParams) error {
			if params.UserID == TestUserID && params.Password == TestUserPassword {
				return nil
			}
			if params.UserID == TestUserID2 && params.Password == TestUserPassword2 {
				return nil
			}
			return tcmrsv.ErrAuthenticationFailed
		},
	}
}

// GetIkebukuroRoom returns the first room found for Ikebukuro campus
func GetIkebukuroRoom(ctx context.Context, t *testing.T, roomRepo *room_repo.Repository) entity.Room {
	rooms := roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{})
	require.NotEmpty(t, rooms)

	for _, room := range rooms {
		if room.CampusType == enum.CampusTypeIkebukuro {
			return room
		}
	}

	t.Fatal("No Ikebukuro room found")
	return entity.Room{}
}

// GetRoomByCampusType returns the first room found for the specified campus type
func GetRoomByCampusType(ctx context.Context, t *testing.T, roomRepo *room_repo.Repository, campusType enum.CampusType) entity.Room {
	rooms := roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{})
	require.NotEmpty(t, rooms)

	for _, room := range rooms {
		if room.CampusType == campusType {
			return room
		}
	}

	t.Fatalf("No room found for campus type %v", campusType)
	return entity.Room{}
}

// CreateTestReservation creates a test reservation with default parameters
func CreateTestReservation(
	ctx context.Context,
	t *testing.T,
	rsvUC *rsv_uc.Usecase,
	userID string,
	room entity.Room,
	date time.Time,
	fromHour, fromMinute, toHour, toMinute int,
) *entity.Reservation {
	output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
		UserID:     userID,
		CampusType: room.CampusType,
		Date:       date,
		FromHour:   fromHour,
		FromMinute: fromMinute,
		ToHour:     toHour,
		ToMinute:   toMinute,
		RoomID:     room.ID,
		BookerName: nil,
	})
	require.NoError(t, err)
	require.Len(t, output.Reservations, 1)
	return &output.Reservations[0]
}

// CreateDefaultTestReservation creates a test reservation with commonly used default values
func CreateDefaultTestReservation(
	ctx context.Context,
	t *testing.T,
	rsvUC *rsv_uc.Usecase,
	userID string,
	room entity.Room,
) *entity.Reservation {
	return CreateTestReservation(
		ctx,
		t,
		rsvUC,
		userID,
		room,
		time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()),
		9, 30,  // from 9:30
		12, 0,  // to 12:00
	)
}

// TestReservationParams holds parameters for creating a test reservation
type TestReservationParams struct {
	UserID     string
	Room       entity.Room
	Date       time.Time
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
	BookerName *string
}

// CreateTestReservationWithParams creates a test reservation with custom parameters
func CreateTestReservationWithParams(
	ctx context.Context,
	t *testing.T,
	rsvUC *rsv_uc.Usecase,
	params TestReservationParams,
) *entity.Reservation {
	output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
		UserID:     params.UserID,
		CampusType: params.Room.CampusType,
		Date:       params.Date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
		RoomID:     params.Room.ID,
		BookerName: params.BookerName,
	})
	require.NoError(t, err)
	require.Len(t, output.Reservations, 1)
	return &output.Reservations[0]
}

// SetupReservationTestDependencies sets up common dependencies for reservation tests
type ReservationTestDeps struct {
	RoomRepo *room_repo.Repository
	RsvRepo  *rsv_repo.Repository
	UserRepo *user_repo.Repository
	RsvUC    *rsv_uc.Usecase
}

func SetupReservationTestDependencies(db database.Execer) *ReservationTestDeps {
	roomRepo := room_repo.NewRepository(tcmrsv.New())
	rsvRepo := rsv_repo.NewRepository(db)
	userRepo := user_repo.NewRepository(db)
	rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

	return &ReservationTestDeps{
		RoomRepo: roomRepo,
		RsvRepo:  rsvRepo,
		UserRepo: userRepo,
		RsvUC:    rsvUC,
	}
}

// GetTestDate returns a commonly used test date
func GetTestDate() time.Time {
	return time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST())
}

// GetTestDateTime returns a commonly used test date with time
func GetTestDateTime() time.Time {
	return time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST())
}