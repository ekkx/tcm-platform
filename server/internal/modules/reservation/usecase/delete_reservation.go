package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *Usecase) DeleteReservation(ctx context.Context, params *input.DeleteReservation) error {
	rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ReservationID)
	if err != nil {
		return err
	}

	// 予約者からのリクエストであるかを確認
	if params.Actor.Role == actor.RoleUser {
		if params.Actor.ID != rsv.UserID {
			return errs.ErrNotYourReservation
		}
	}

	// 予約を削除
	err = uc.rsvRepo.DeleteReservationByID(ctx, params.ReservationID)
	if err != nil {
		return err
	}

	return nil
}
