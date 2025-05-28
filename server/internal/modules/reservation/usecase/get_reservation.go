package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/core/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
)

func (uc *Usecase) GetReservation(ctx context.Context, params *input.GetReservation) (*output.GetReservation, error) {
	// 予約情報を取得
	rsv, err := uc.rsvRepo.GetReservationByID(ctx, int64(params.ReservationID))
	if err != nil {
		return nil, err
	}

	// ユーザーがリクエストを行う場合、予約の所有者と一致するか確認
	if (params.Actor.Role == actor.RoleUser) && (params.Actor.ID != rsv.UserID) {
		return nil, apperrors.ErrNotYourReservation
	}

	return output.NewGetReservation(rsv), nil
}
