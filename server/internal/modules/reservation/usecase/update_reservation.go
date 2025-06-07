package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *Usecase) UpdateReservation(ctx context.Context, params *input.UpdateReservation) (*output.UpdateReservation, error) {
	if err := params.Validate(); err != nil {
		return nil, errs.ErrInvalidArgument.WithCause(err)
	}

	// 日付の時間部分を0に正規化
	params.Date = time.Date(params.Date.Year(), params.Date.Month(), params.Date.Day(), 0, 0, 0, 0, params.Date.Location())

	// 予約更新を行う権限があるか確認
	if params.Actor.Role == actor.RoleUser {
		rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ID)
		if err != nil {
			return nil, err
		}

		if rsv == nil {
			return nil, errs.ErrReservationNotFound
		}

		if rsv.UserID != params.Actor.ID {
			return nil, errs.ErrNotYourReservation
		}
	}

	// 予約時間と練習室が被っていないか確認
	originalRsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if originalRsv == nil {
		return nil, errs.ErrReservationNotFound
	}

	// 同じ日付の予約を取得
	reservations, err := uc.rsvRepo.GetReservationsByDate(ctx, &repository.GetReservationsByDate{
		Date: params.Date,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	// 自分自身の予約を除外して、他の予約と時間が重複していないか確認
	for _, rsv := range reservations {
		// 自分自身の予約はスキップ
		if rsv.ID == params.ID {
			continue
		}

		// 同じ部屋の予約のみチェック
		if rsv.RoomID != params.RoomID {
			continue
		}

		// 時間の重複チェック
		newStartMinutes := params.FromHour*60 + params.FromMinute
		newEndMinutes := params.ToHour*60 + params.ToMinute
		existingStartMinutes := rsv.FromHour*60 + rsv.FromMinute
		existingEndMinutes := rsv.ToHour*60 + rsv.ToMinute

		// 時間が重複している場合
		if newStartMinutes < existingEndMinutes && newEndMinutes > existingStartMinutes {
			return nil, errs.ErrReservationConflict
		}
	}

	err = uc.rsvRepo.UpdateReservationByID(ctx, &repository.UpdateReservationByIDArgs{
		ExternalID:    params.ExternalID,
		CampusType:    &params.CampusType,
		RoomID:        &params.RoomID,
		Date:          &params.Date,
		FromHour:      &params.FromHour,
		FromMinute:    &params.FromMinute,
		ToHour:        &params.ToHour,
		ToMinute:      &params.ToMinute,
		BookerName:    params.BookerName,
		ReservationID: params.ID,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ID)
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	if rsv == nil {
		return nil, errs.ErrReservationNotFound
	}

	return output.NewUpdateReservation(*rsv), nil
}
