package adapter

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type QueryAdapter struct {
	reservationRepo repository.Repository
}

func NewQueryAdapter(reservationRepo repository.Repository) gateway.ReservationQuery {
	return &QueryAdapter{
		reservationRepo: reservationRepo,
	}
}

func (q *QueryAdapter) GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error) {
	return q.reservationRepo.GetReservationByID(ctx, reservationID)
}

func (q *QueryAdapter) ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error) {
	return q.reservationRepo.ListReservationsByIDs(ctx, reservationIDs)
}
