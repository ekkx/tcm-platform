package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
)

type GetUserReservationsArgs struct {
	UserID string    `json:"user_id"`
	Date   time.Time `json:"date"`
}

func (r *Repository) GetUserReservations(ctx context.Context, args *GetUserReservationsArgs) ([]entity.Reservation, error) {
	rows, err := r.db.Query(ctx, `
        SELECT
            reservations.*
        FROM
            reservations
        WHERE
            reservations.user_id = $1
            AND reservations.date >= $2
    `, args.UserID, args.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.Reservation{}
	for rows.Next() {
		var rsv entity.Reservation
		if err := rows.Scan( // TODO: 共通化できるかも
			&rsv.ID, &rsv.ExternalID, &rsv.UserID, &rsv.CampusType, &rsv.RoomID, &rsv.Date,
			&rsv.FromHour, &rsv.FromMinute, &rsv.ToHour, &rsv.ToMinute, &rsv.BookerName, &rsv.CreatedAt,
		); err != nil {
			return nil, apperrors.ErrInternal.WithCause(err)
		}
		items = append(items, rsv)
	}
	if err := rows.Err(); err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	return items, nil
}
