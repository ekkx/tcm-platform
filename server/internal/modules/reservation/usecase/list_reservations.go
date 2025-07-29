package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/service"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

func (uc *UseCaseImpl) ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	rsvs, err := uc.reservationService.ListUserReservations(ctx, &service.ListUserReservationsParams{
		UserID: input.Actor.ID,
		Date:   ymd.Today(),
	})
	if err != nil {
		return nil, err
	}

	return NewListReservationsOutput(rsvs), nil
}
