package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/apperrors"
)

type CheckReservationConflictArgs struct {
	RoomID     string    `json:"room_id"`
	Date       time.Time `json:"date"`
	FromHour   int       `json:"from_hour"`
	FromMinute int       `json:"from_minute"`
	ToHour     int       `json:"to_hour"`
	ToMinute   int       `json:"to_minute"`
}

func (r *Repository) CheckReservationConflict(ctx context.Context, args *CheckReservationConflictArgs) (bool, error) {
	row := r.db.QueryRow(ctx, `
        SELECT EXISTS (
            SELECT
                1
            FROM
                reservations
            WHERE
                room_id = $1
                AND date = $2
                AND (
                    (($3 * 60) + $4) < ((to_hour * 60) + to_minute)
                    AND (($5 * 60) + $6) > ((from_hour * 60) + from_minute)
                )
        ) AS conflict
    `, args.RoomID, args.Date, args.FromHour, args.FromMinute, args.ToHour, args.ToMinute)

	var conflict bool
	err := row.Scan(&conflict)
	if err != nil {
		return false, apperrors.ErrInternal.WithCause(err)
	}

	return conflict, nil
}
