package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/domain"
)

func toReservation(v *domain.Reservation) *api.Reservation {
	if v == nil {
		return nil
	}

	return &api.Reservation{
		BookerName: v.BookerName,
		CampusCode: api.ReservationCampusCode(v.Campus),
		CreatedAt:  v.CreatedAt,
		Date:       v.Date,
		ExternalId: v.ExternalID,
		FromHour:   v.FromHour,
		FromMinute: v.FromMinute,
		Id:         v.ID,
		RoomId:     v.RoomID,
		ToHour:     v.ToHour,
		ToMinute:   v.ToMinute,
	}
}

func toReservationList(v *[]domain.Reservation) []api.Reservation {
	if v == nil {
		return []api.Reservation{}
	}

	list := make([]api.Reservation, 0, len(*v))
	for _, rsv := range *v {
		if converted := toReservation(&rsv); converted != nil {
			list = append(list, *converted)
		}
	}
	return list
}
