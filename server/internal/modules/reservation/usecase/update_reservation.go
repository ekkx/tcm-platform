package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
)

func (uc *Usecase) UpdateReservation(ctx context.Context, params *input.UpdateReservation) (*output.UpdateReservation, error) {
	// Validate the input parameters
	// if err := params.Validate(); err != nil {
	//     return nil, err
	// }

	// 予約更新を行う権限があるか確認
	if params.Actor.Role == actor.RoleUser {
		rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ID)
		if err != nil {
			return nil, err
		}

		if rsv.UserID != params.Actor.ID {
			return nil, apperrors.ErrNotYourReservation
		}
	}

	err := uc.rsvRepo.UpdateReservationByID(ctx, &repository.UpdateReservationByIDArgs{
		ExternalID:    params.ExternalID,
		CampusType:    &params.CampusType,
		RoomID:        params.RoomID,
		Date:          &params.Date,
		FromHour:      &params.FromHour,
		FromMinute:    &params.FromMinute,
		ToHour:        &params.ToHour,
		ToMinute:      &params.ToMinute,
		BookerName:    params.BookerName,
		ReservationID: params.ID,
	})
	if err != nil {
		return nil, err
	}

	rsv, err := uc.rsvRepo.GetReservationByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return output.NewUpdateReservation(rsv), nil
}
