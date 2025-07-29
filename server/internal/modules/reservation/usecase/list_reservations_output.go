package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type ListReservationsOutput struct {
	Reservations []*entity.Reservation
}

func NewListReservationsOutput(reservations []*entity.Reservation) *ListReservationsOutput {
	return &ListReservationsOutput{
		Reservations: reservations,
	}
}

func (st *ListReservationsOutput) ToResponse() *connect.Response[reservationv1.ListReservationsResponse] {
	return connect.NewResponse(&reservationv1.ListReservationsResponse{
		Reservations: presenter.ToReservationList(st.Reservations),
	})
}
