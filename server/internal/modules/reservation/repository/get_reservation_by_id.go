package repository

import (
	"context"
	"errors"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetReservationByID(ctx context.Context, id int) (*entity.Reservation, error) {
	row := r.db.QueryRow(ctx, `
        SELECT
            reservations.*
        FROM
            reservations
        WHERE
            reservations.id = $1
    `, id,
	)

	var rsv entity.Reservation
	err := row.Scan(
		&rsv.ID, &rsv.ExternalID, &rsv.UserID, &rsv.CampusType, &rsv.RoomID, &rsv.Date,
		&rsv.FromHour, &rsv.FromMinute, &rsv.ToHour, &rsv.ToMinute, &rsv.BookerName, &rsv.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	rsv.Date = rsv.Date.In(utils.JST())

	return &rsv, nil
}
