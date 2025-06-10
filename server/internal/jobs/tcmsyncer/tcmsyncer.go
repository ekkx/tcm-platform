package tcmsyncer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	userRepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
)

// TCMRSVClient interface for testing
type TCMRSVClient interface {
	Login(params *tcmrsv.LoginParams) error
	Reserve(params *tcmrsv.ReserveParams) error
	GetMyReservations() ([]tcmrsv.Reservation, error)
}

// tcmrsvClientWrapper wraps the actual tcmrsv.Client
type tcmrsvClientWrapper struct {
	client *tcmrsv.Client
}

func (w *tcmrsvClientWrapper) Login(params *tcmrsv.LoginParams) error {
	return w.client.Login(params)
}

func (w *tcmrsvClientWrapper) Reserve(params *tcmrsv.ReserveParams) error {
	return w.client.Reserve(params)
}

func (w *tcmrsvClientWrapper) GetMyReservations() ([]tcmrsv.Reservation, error) {
	return w.client.GetMyReservations()
}

type SyncReservationsJob struct {
	db         database.Execer
	reservRepo ReservationRepository
	userRepo   UserRepository
	aesKey     []byte
	// For testing
	tcmClientFactory func() TCMRSVClient
}

// ReservationRepository interface
type ReservationRepository interface {
	GetReservationsByDate(ctx context.Context, args *repository.GetReservationsByDate) ([]entity.Reservation, error)
	UpdateReservationByID(ctx context.Context, args *repository.UpdateReservationByIDArgs) error
}

// UserRepository interface
type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}

func NewSyncReservationsJob(db database.Execer, aesKey []byte) *SyncReservationsJob {
	return &SyncReservationsJob{
		db:         db,
		reservRepo: repository.NewRepository(db),
		userRepo:   userRepo.NewRepository(db),
		aesKey:     aesKey,
		tcmClientFactory: func() TCMRSVClient {
			client := tcmrsv.New()
			return &tcmrsvClientWrapper{client: client}
		},
	}
}

// NewSyncReservationsJobWithDeps creates a new job with injected dependencies (for testing)
func NewSyncReservationsJobWithDeps(db database.Execer, reservRepo ReservationRepository, userRepo UserRepository, aesKey []byte) *SyncReservationsJob {
	return &SyncReservationsJob{
		db:         db,
		reservRepo: reservRepo,
		userRepo:   userRepo,
		aesKey:     aesKey,
		tcmClientFactory: func() TCMRSVClient {
			client := tcmrsv.New()
			return &tcmrsvClientWrapper{client: client}
		},
	}
}

// SetTCMClientFactory sets the client factory (for testing)
func (j *SyncReservationsJob) SetTCMClientFactory(factory func() TCMRSVClient) {
	j.tcmClientFactory = factory
}

func (j *SyncReservationsJob) Execute(ctx context.Context) error {
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	reservations, err := j.reservRepo.GetReservationsByDate(ctx, &repository.GetReservationsByDate{
		Date: today,
	})
	if err != nil {
		return fmt.Errorf("failed to get reservations: %w", err)
	}

	var todayReservations []entity.Reservation
	for _, r := range reservations {
		if r.Date.After(today) && r.Date.Before(tomorrow) {
			todayReservations = append(todayReservations, r)
		}
	}

	log.Printf("Found %d reservations for %s", len(todayReservations), today.Format("2006-01-02"))

	// Group reservations by user
	userReservations := make(map[string][]entity.Reservation)
	for _, reservation := range todayReservations {
		if reservation.ExternalID != nil {
			log.Printf("Reservation %d already has external ID %s, skipping", reservation.ID, *reservation.ExternalID)
			continue
		}
		userReservations[reservation.UserID] = append(userReservations[reservation.UserID], reservation)
	}

	// Process reservations by user
	for userID, userResvs := range userReservations {
		// Login once per user
		user, err := j.userRepo.GetUserByID(ctx, userID)
		if err != nil {
			log.Printf("Failed to get user %s: %v", userID, err)
			continue
		}

		password, err := cryptohelper.DecryptAES(user.EncryptedPassword, j.aesKey)
		if err != nil {
			log.Printf("Failed to decrypt password for user %s: %v", userID, err)
			continue
		}

		client := j.tcmClientFactory()
		loginParams := &tcmrsv.LoginParams{
			UserID:   user.ID,
			Password: password,
		}

		if err := client.Login(loginParams); err != nil {
			log.Printf("Failed to login as user %s: %v", userID, err)
			continue
		}

		// Create all reservations for this user
		type pendingReservation struct {
			reservation entity.Reservation
			campus      tcmrsv.Campus
		}
		var pendingReservations []pendingReservation

		for _, reservation := range userResvs {
			bookerName := reservation.BookerName
			if bookerName == nil {
				defaultName := "予約"
				bookerName = &defaultName
			}

			// Convert CampusType enum to tcmrsv.Campus string
			var campus tcmrsv.Campus
			switch reservation.CampusType {
			case enum.CampusTypeNakameguro:
				campus = tcmrsv.CampusNakameguro
			case enum.CampusTypeIkebukuro:
				campus = tcmrsv.CampusIkebukuro
			default:
				log.Printf("Unknown campus type %d for reservation %d", reservation.CampusType, reservation.ID)
				continue
			}

			err = client.Reserve(&tcmrsv.ReserveParams{
				Campus:     campus,
				RoomID:     reservation.RoomID,
				Date:       reservation.Date,
				FromHour:   reservation.FromHour,
				FromMinute: reservation.FromMinute,
				ToHour:     reservation.ToHour,
				ToMinute:   reservation.ToMinute,
			})
			if err != nil {
				log.Printf("Failed to create external reservation for reservation %d: %v", reservation.ID, err)
				continue
			}

			pendingReservations = append(pendingReservations, pendingReservation{
				reservation: reservation,
				campus:      campus,
			})
		}

		// Get all reservations once after creating all reservations for this user
		if len(pendingReservations) > 0 {
			myReservations, err := client.GetMyReservations()
			if err != nil {
				log.Printf("Failed to get my reservations for user %s: %v", userID, err)
				continue
			}

			// Match each pending reservation with external IDs
			for _, pending := range pendingReservations {
				var externalID string
				targetDate := pending.reservation.Date.Format("2006年01月02日")
				targetTimeRange := fmt.Sprintf("%02d:%02d-%02d:%02d", 
					pending.reservation.FromHour, pending.reservation.FromMinute, 
					pending.reservation.ToHour, pending.reservation.ToMinute)

				for _, myRes := range myReservations {
					// Check if date, time, room, and campus match
					if myRes.Date == targetDate &&
					   myRes.TimeRange == targetTimeRange &&
					   myRes.RoomName == pending.reservation.RoomID &&
					   myRes.Campus == pending.campus {
						externalID = myRes.ID
						break
					}
				}

				if externalID == "" {
					log.Printf("Could not find external ID for reservation %d (date: %s, time: %s, room: %s, campus: %s)",
						pending.reservation.ID, targetDate, targetTimeRange, pending.reservation.RoomID, pending.campus)
					continue
				}

				err = j.reservRepo.UpdateReservationByID(ctx, &repository.UpdateReservationByIDArgs{
					ReservationID: pending.reservation.ID,
					ExternalID:    &externalID,
				})
				if err != nil {
					log.Printf("Failed to update external ID for reservation %d: %v", pending.reservation.ID, err)
					continue
				}

				log.Printf("Successfully created external reservation %s for reservation %d", externalID, pending.reservation.ID)
			}
		}
	}

	return nil
}

func (j *SyncReservationsJob) RunAt12PM(ctx context.Context) {
	location := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(location)

	next12PM := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, location)
	if now.After(next12PM) {
		next12PM = next12PM.Add(24 * time.Hour)
	}

	durationUntil12PM := next12PM.Sub(now)
	log.Printf("Next sync will run at %s (in %v)", next12PM.Format("2006-01-02 15:04:05"), durationUntil12PM)

	timer := time.NewTimer(durationUntil12PM)
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			j.runSync(ctx)
		case <-ticker.C:
			j.runSync(ctx)
		}
	}
}

func (j *SyncReservationsJob) runSync(ctx context.Context) {
	log.Println("Starting reservation sync...")
	if err := j.Execute(ctx); err != nil {
		log.Printf("Sync failed: %v", err)
	} else {
		log.Println("Sync completed successfully")
	}
}
