package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UseCaseImpl) DeleteReservation(ctx context.Context, input *DeleteReservationInput) (*DeleteReservationOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	rsv, err := uc.reservationService.GetReservationByID(ctx, input.ReservationID)
	if err != nil {
		return nil, err
	}

	if rsv == nil {
		return nil, errs.ErrReservationNotFound
	}

	// 予約を取ったユーザー本人、またはそのユーザーのマスターユーザーのみ削除可能
	if rsv.User.ID != input.Actor.ID {
		if rsv.User.MasterUser == nil || rsv.User.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only delete your own reservations")
		}
	}

	if err := uc.reservationRepo.DeleteReservationByID(ctx, input.ReservationID); err != nil {
		return nil, err
	}

	return NewDeleteReservationOutput(), nil
}
