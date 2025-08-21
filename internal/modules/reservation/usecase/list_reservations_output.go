package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
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
