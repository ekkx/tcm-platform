package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/core/enum"
)

type UpdateReservationByIDArgs struct {
	ExternalID    *string
	CampusType    *enum.CampusType
	RoomID        *string
	Date          *time.Time
	FromHour      *int32
	FromMinute    *int32
	ToHour        *int32
	ToMinute      *int32
	BookerName    *string
	ReservationID int64
}

func (r *Repository) UpdateReservationByID(ctx context.Context, args *UpdateReservationByIDArgs) error {
	row := r.db.QueryRow(ctx, `
        UPDATE
            reservations
        SET
            external_id = COALESCE($1, custom_id),
            campus_type = COALESCE($2, campus_type),
            room_id = COALESCE($3, room_id),
            date = COALESCE($4, date),
            from_hour = COALESCE($5, from_hour),
            from_minute = COALESCE($6, from_minute),
            to_hour = COALESCE($7, to_hour),
            to_minute = COALESCE($8, to_minute),
            booker_name = COALESCE($9, booker_name)
        WHERE
            id = $10
        RETURNING 1
    `, args.ExternalID, args.CampusType, args.RoomID, args.Date,
		args.FromHour, args.FromMinute, args.ToHour, args.ToMinute, args.BookerName, args.ReservationID,
	)

	var n int
	err := row.Scan(&n)
	if err != nil {
		return apperrors.ErrInternal.WithCause(err)
	}

	return nil
}
