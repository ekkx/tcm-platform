package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

func (u *Usecase) GetUserReservations(ctx context.Context, params *input.GetUserReservations) (*output.GetMyReservations, error) {
	rsvs, err := u.rsvRepo.GetMyReservations(ctx, &repository.GetMyReservationsArgs{
		UserID: params.UserID,
		Date:   params.Date,
	})
	if err != nil {
		return nil, err
	}

	return output.NewGetMyReservations(rsvs), nil
}
