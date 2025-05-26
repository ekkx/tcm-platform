package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

type GetMyReservationsInput struct {
	UserID string
	Date   time.Time
}

type GetMyReservationsOutput struct {
	Reservations []entity.Reservation
}

func (u *Usecase) GetMyReservations(ctx context.Context, input *GetMyReservationsInput) (*GetMyReservationsOutput, error) {
	rsvs, err := u.rsvrepo.GetMyReservations(ctx, &repository.GetMyReservationsArgs{
		UserID: input.UserID,
		Date:   input.Date,
	})
	if err != nil {
		return nil, err
	}

	return &GetMyReservationsOutput{
		Reservations: rsvs,
	}, nil
}
