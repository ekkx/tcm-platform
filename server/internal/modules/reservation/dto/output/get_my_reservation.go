package output

import "github.com/ekkx/tcmrsv-web/server/internal/core/entity"

type GetMyReservations struct {
	Reservations []entity.Reservation
}

func NewGetMyReservations(reservations []entity.Reservation) *GetMyReservations {
	return &GetMyReservations{
		Reservations: reservations,
	}
}
