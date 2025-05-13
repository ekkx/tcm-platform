package reservation

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/jackc/pgx/v5"
)

type DeleteReservationInput struct {
	ReservationID int
}

func (uc *ReservationUsecaseImpl) DeleteReservation(ctx context.Context, input *DeleteReservationInput) error {
	_, err := uc.querier.DeleteReservationByID(ctx, int32(input.ReservationID))
	if err != nil {
		if err == pgx.ErrNoRows {
			return apperrors.ErrReservationNotFound
		}
		return err
	}

	return nil
}
