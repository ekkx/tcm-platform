package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) UpdateReservationNote(ctx context.Context, input *UpdateReservationNoteInput) (*UpdateReservationNoteOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	v, err := uc.reservationAsm.Build(ctx, input.ReservationID)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return nil, errs.ErrReservationNotFound
	}

	if v.UserView.User.ID != input.Actor.ID {
		if v.UserView.MasterUser == nil || v.UserView.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only update your own reservations")
		}
	}

	err = uc.reservationRepo.UpdateReservationNote(ctx, &repository.UpdateReservationNoteParams{
		ID:   input.ReservationID,
		Note: input.Note,
	})
	if err != nil {
		return nil, err
	}

	updated, err := uc.reservationAsm.Build(ctx, input.ReservationID)
	if err != nil {
		return nil, err
	}

	return NewUpdateReservationNoteOutput(*updated), nil
}
