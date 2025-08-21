package gateway

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type ReservationQuery interface {
	GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error)
	ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
}
