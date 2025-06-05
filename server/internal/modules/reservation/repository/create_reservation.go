package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

type CreateReservationArgs struct {
	UserID     string          `json:"user_id"`
	CampusType enum.CampusType `json:"campus_type"`
	RoomID     string          `json:"room_id"`
	Date       time.Time       `json:"date"`
	FromHour   int             `json:"from_hour"`
	FromMinute int             `json:"from_minute"`
	ToHour     int             `json:"to_hour"`
	ToMinute   int             `json:"to_minute"`
	BookerName *string         `json:"booker_name"`
}

func (r *Repository) CreateReservation(ctx context.Context, args *CreateReservationArgs) (entity.Reservation, error) {
	row := r.db.QueryRow(ctx, `
        INSERT INTO
            reservations (
                user_id, campus_type, room_id, date,
                from_hour, from_minute, to_hour, to_minute, booker_name
            )
        VALUES
            (
                $1, $2, $3, $4,
                $5, $6, $7, $8, $9
            )
        RETURNING
            reservations.*`,
		args.UserID, args.CampusType, args.RoomID, args.Date,
		args.FromHour, args.FromMinute, args.ToHour, args.ToMinute, args.BookerName,
	)

	var rsv entity.Reservation
	err := row.Scan(
		&rsv.ID, &rsv.ExternalID, &rsv.UserID, &rsv.CampusType, &rsv.RoomID, &rsv.Date,
		&rsv.FromHour, &rsv.FromMinute, &rsv.ToHour, &rsv.ToMinute, &rsv.BookerName, &rsv.CreatedAt,
	)
	if err != nil {
		return entity.Reservation{}, errs.ErrInternal.WithCause(err)
	}

	rsv.Date = rsv.Date.In(utils.JST())

	return rsv, nil
}
