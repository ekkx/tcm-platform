package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
)

func (uc *UseCaseImpl) DeleteReservation(ctx context.Context, input *DeleteReservationInput) (*DeleteReservationOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	rsv, err := uc.reservationRepo.GetReservationByID(ctx, input.ReservationID)
	if err != nil {
		return nil, err
	}

	if rsv == nil {
		return nil, errs.ErrReservationNotFound
	}

	user, err := uc.userQuery.GetUserByID(ctx, rsv.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.ErrUserNotFound.WithMessage("user associated with the reservation not found")
	}

	// 予約を取ったユーザー本人、またはそのユーザーのマスターユーザーのみ削除可能
	if user.ID != input.Actor.ID {
		if user.MasterUserID == nil || *user.MasterUserID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only delete your own reservations")
		}
	}

	if err := uc.reservationRepo.DeleteReservationByID(ctx, input.ReservationID); err != nil {
		return nil, err
	}

	return NewDeleteReservationOutput(), nil
}
