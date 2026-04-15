package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcmrsv"
)

func (uc *UseCaseImpl) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// サブスクリプションのチェック
	sub, err := uc.subRepo.GetSubscriptionByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}
	if sub == nil || (!sub.IsUnlimited() && sub.Status != "active") {
		return nil, errs.ErrNoActiveSubscription
	}

	// 利用時間の制限チェック（unlimited以外）
	if !sub.IsUnlimited() && sub.MonthlyHours != nil {
		usedMinutes, err := uc.subRepo.GetUsedMinutesByUserID(ctx, input.Actor.ID)
		if err != nil {
			return nil, err
		}
		newMinutes := int32((input.ToHour*60 + input.ToMinute) - (input.FromHour*60 + input.FromMinute))
		limitMinutes := int32(*sub.MonthlyHours * 60)
		if usedMinutes+newMinutes > limitMinutes {
			return nil, errs.ErrUsageLimitExceeded
		}
	}

	// ルームの存在チェック
	rooms := tcmrsv.New().GetRoomsFiltered(tcmrsv.GetRoomsFilteredParams{
		ID: &input.RoomID,
	})
	if len(rooms) == 0 {
		return nil, errs.ErrRoomNotFound
	}

	isConflicted, err := uc.reservationRepo.IsReservationConflicted(ctx, &repository.IsReservationConflictedParams{
		RoomID:     input.RoomID,
		Date:       input.Date,
		FromHour:   input.FromHour,
		FromMinute: input.FromMinute,
		ToHour:     input.ToHour,
		ToMinute:   input.ToMinute,
	})
	if err != nil {
		return nil, err
	}
	if isConflicted {
		return nil, errs.ErrReservationConflict
	}

	reservationID, err := uc.reservationRepo.CreateReservation(ctx, &repository.CreateReservationParams{
		UserID:     input.Actor.ID,
		CampusType: input.CampusType,
		RoomID:     input.RoomID,
		Date:       input.Date,
		FromHour:   input.FromHour,
		FromMinute: input.FromMinute,
		ToHour:     input.ToHour,
		ToMinute:   input.ToMinute,
		Note:       input.Note,
	})
	if err != nil {
		return nil, err
	}

	reservation, err := uc.reservationAsm.Build(ctx, *reservationID)
	if err != nil {
		return nil, err
	}

	return NewCreateReservationOutput(*reservation), nil
}
