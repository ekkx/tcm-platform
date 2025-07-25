package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

func (uc *UseCaseImpl) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// TODO: コンフリクトチェック

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
