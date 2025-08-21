package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/pkg/ymd"
)

func (uc *UseCaseImpl) ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	rsvIDs, err := uc.reservationRepo.ListUserReservationIDs(ctx, &repository.ListUserReservationIDsParams{
		UserID: input.Actor.ID,
		Date:   ymd.Today(),
	})
	if err != nil {
		return nil, err
	}

	if len(rsvIDs) == 0 {
		return NewListReservationsOutput(nil), nil
	}

	rsvs, err := uc.reservationAsm.BuildList(ctx, rsvIDs)
	if err != nil {
		return nil, err
	}

	return NewListReservationsOutput(rsvs), nil
}
