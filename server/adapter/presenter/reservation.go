package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
)

func GetMyReservations(output *reservation.GetMyReservationsOutput) *api.ReservationList {
	return &api.ReservationList{
		Reservations: toReservationList(&output.Reservations),
	}
}

func CreateReservation(output *reservation.CreateReservationOutput) *api.Reservation {
	return toReservation(&output.Reservation)
}
