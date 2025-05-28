package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/core/apperrors"
)

func (r *Repository) DeleteReservationByID(ctx context.Context, reservationID int64) error {
	row := r.db.QueryRow(ctx, `
        DELETE FROM
            reservations
        WHERE
            id = $1
        RETURNING 1
    `, reservationID)

	var n int
	err := row.Scan(&n)
	if err != nil {
		return apperrors.ErrInternal.WithCause(err)
	}

	return nil
}
