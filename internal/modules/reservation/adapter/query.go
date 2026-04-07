package adapter

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
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

func (q *QueryAdapter) ListReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error) {
	if date.IsZero() || !date.IsValid() {
		return nil, nil
	}
	return q.reservationRepo.ListReservationIDsByDate(ctx, date)
}

func (q *QueryAdapter) ListPendingReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error) {
	if date.IsZero() || !date.IsValid() {
		return nil, nil
	}
	return q.reservationRepo.ListPendingReservationIDsByDate(ctx, date)
}
