package reservation

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/adapter/db/mapper"
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
)

type GetMyReservationsInput struct {
	UserID   string
	Password string

	Campus    *domain.Campus
	PianoType *domain.PianoType
	Date      *time.Time
}

type GetMyReservationsOutput struct {
	Reservations []domain.Reservation
}

func (uc *ReservationUsecaseImpl) GetMyReservations(ctx context.Context, input *GetMyReservationsInput) (*GetMyReservationsOutput, error) {
	var date time.Time
	if input.Date != nil {
		date = time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, tcmrsv.JST())
	} else {
		now := time.Now()
		date = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, tcmrsv.JST())
	}

	rsvs, err := uc.querier.GetMyReservations(ctx, db.GetMyReservationsParams{
		UserID: input.UserID,
		Date:   date,
		// TODO: キャンパスコードとピアノの種類で絞り込みできるようにする
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
