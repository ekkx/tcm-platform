package gateway

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type ReservationQuery interface {
	GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error)
	ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
	ListReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error)
	ListPendingReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error)
}

type ReservationCommand interface {
	UpdateReservationStatus(ctx context.Context, params *UpdateReservationStatusCommand) error
}

type UpdateReservationStatusCommand struct {
	ID             ulid.ULID
	Status         valueobject.ReservationStatus
	OfficialSiteID *string
}
