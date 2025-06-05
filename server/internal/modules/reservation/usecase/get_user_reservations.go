package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (u *Usecase) GetUserReservations(ctx context.Context, params *input.GetUserReservations) (*output.GetMyReservations, error) {
	rsvs, err := u.rsvRepo.GetUserReservations(ctx, &repository.GetUserReservationsArgs{
		UserID: params.UserID,
		Date:   params.Date,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	return output.NewGetMyReservations(rsvs), nil
}
