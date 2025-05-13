package reservation

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/adapter/db/mapper"
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

type GetMyReservationsInput struct {
	UserID   string
	Password string
}

type GetMyReservationsOutput struct {
	Reservations []domain.Reservation
}

func (uc *ReservationUsecaseImpl) GetMyReservations(ctx context.Context, input *GetMyReservationsInput) (*GetMyReservationsOutput, error) {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, utils.JST())

	rsvs, err := uc.querier.GetMyReservations(ctx, db.GetMyReservationsParams{
		UserID: input.UserID,
		Date:   date,
	})
	if err != nil {
		return nil, err
	}

	var domainRsvs = make([]domain.Reservation, 0, len(rsvs))
	for _, rsv := range rsvs {
		domainRsvs = append(domainRsvs, mapper.ToReservation(rsv))
	}

	return &GetMyReservationsOutput{
		Reservations: domainRsvs,
	}, nil
}
