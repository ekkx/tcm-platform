package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	reservationv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
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
