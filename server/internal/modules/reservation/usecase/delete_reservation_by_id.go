package usecase

import (
	"context"
)

type DeleteReservationByIDInput struct {
	ReservationID int
}

func (u *Usecase) DeleteReservationByID(ctx context.Context, input *DeleteReservationByIDInput) error {
	_, err := u.rsvrepo.DeleteReservationByID(ctx, input.ReservationID)
	if err != nil {
		return err
	}

	return nil
}
