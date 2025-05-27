package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
)

func (u *Usecase) DeleteReservationByID(ctx context.Context, params *input.DeleteReservationByID) error {
	err := u.rsvrepo.DeleteReservationByID(ctx, params.ReservationID)
	if err != nil {
		return err
	}

	return nil
}
