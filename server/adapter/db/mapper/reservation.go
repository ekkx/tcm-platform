package mapper

import (
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
)

func ToReservation(rsv db.Reservation) domain.Reservation {
	return domain.Reservation{
		ID:         int(rsv.ID),
		ExternalID: rsv.ExternalID,
		Campus:     domain.Campus(rsv.Campus),
		RoomID:     rsv.RoomID,
		Date:       rsv.Date.Time,
		FromHour:   int(rsv.FromHour),
		FromMinute: int(rsv.FromMinute),
		ToHour:     int(rsv.ToHour),
		ToMinute:   int(rsv.ToMinute),
		BookerName: rsv.BookerName,
		CreatedAt:  rsv.CreatedAt.Time,
	}
}
