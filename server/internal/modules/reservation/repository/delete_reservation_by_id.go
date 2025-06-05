package repository

import (
	"context"
)

func (r *Repository) DeleteReservationByID(ctx context.Context, reservationID int) error {
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
		return err
	}

	return nil
}
