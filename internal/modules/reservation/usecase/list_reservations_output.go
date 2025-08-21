package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	reservationv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type ListReservationsOutput struct {
	Reservations []*assemble.ReservationView
}

func NewListReservationsOutput(v []*assemble.ReservationView) *ListReservationsOutput {
	return &ListReservationsOutput{
		Reservations: v,
	}
}

func (st *ListReservationsOutput) ToResponse() *connect.Response[reservationv1.ListReservationsResponse] {
	return connect.NewResponse(&reservationv1.ListReservationsResponse{
		Reservations: presenter.ToReservationList(st.Reservations),
	})
}
