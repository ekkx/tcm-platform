package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UsecaseImpl) GetReservation(ctx context.Context, params *input.GetReservation) (*output.GetReservation, error) {
	// 予約情報を取得
	rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ReservationID)
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	if rsv == nil {
		return nil, errs.ErrReservationNotFound
	}

	// ユーザーがリクエストを行う場合、予約の所有者と一致するか確認
	if (params.Actor.Role == actor.RoleUser) && (params.Actor.ID != rsv.UserID) {
		return nil, errs.ErrNotYourReservation
	}

	return output.NewGetReservation(*rsv), nil
}
