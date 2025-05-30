package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/apperrors"
)

type GetReservationsByDate struct {
	Date time.Time `json:"date"`
}

func (r *Repository) GetReservationsByDate(ctx context.Context, args *GetReservationsByDate) ([]entity.Reservation, error) {
	rows, err := r.db.Query(ctx, `
        SELECT
            reservations.*
        FROM
            reservations
        WHERE
            reservations.date = $1
    `, args.Date)
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
