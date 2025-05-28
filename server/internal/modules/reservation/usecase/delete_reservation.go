package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
)

func (uc *Usecase) DeleteReservation(ctx context.Context, params *input.DeleteReservation) error {
	if params.Actor.Role == actor.RoleUser {
		// 予約者からのリクエストであるかを確認
		rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ReservationID)
		if err != nil {
			return err
		}

		if params.Actor.ID != rsv.UserID {
			return apperrors.ErrNotYourReservation
		}
	}

	// 予約を削除
	err := uc.rsvRepo.DeleteReservationByID(ctx, params.ReservationID)
	if err != nil {
		return err
	}

	return nil
}
