package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
)

type GetReservationOutput struct {
	Reservation assemble.ReservationView
}

func NewGetReservationOutput(v assemble.ReservationView) *GetReservationOutput {
	return &GetReservationOutput{
		Reservation: v,
	}
}

func (st *GetReservationOutput) ToResponse() *connect.Response[reservationv1.GetReservationResponse] {
	return connect.NewResponse(&reservationv1.GetReservationResponse{
		Reservation: presenter.ToReservation(&st.Reservation),
	})
}
