package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) GetReservation(ctx context.Context, input *GetReservationInput) (*GetReservationOutput, error) {
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

	// 予約を取ったユーザー本人、またはそのユーザーのマスターユーザーのみ閲覧可能
	if v.UserView.User.ID != input.Actor.ID {
		if v.UserView.MasterUser == nil || v.UserView.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only view your own reservations")
		}
	}

	return NewGetReservationOutput(*v), nil
}
