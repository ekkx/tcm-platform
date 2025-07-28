package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UseCaseImpl) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
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
	})
	if err != nil {
		return nil, err
	}

	reservation, err := uc.reservationService.GetReservationByID(ctx, *reservationID)
	if err != nil {
		return nil, err
	}

	return NewCreateReservationOutput(*reservation), nil
}
