package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
)

type CreateReservationOutput struct {
	Reservation assemble.ReservationView
}

func NewCreateReservationOutput(reservation assemble.ReservationView) *CreateReservationOutput {
	return &CreateReservationOutput{
		Reservation: reservation,
	}
}

func (st *CreateReservationOutput) ToResponse() *connect.Response[reservationv1.CreateReservationResponse] {
	return connect.NewResponse(&reservationv1.CreateReservationResponse{
		Reservation: presenter.ToReservation(&st.Reservation),
	})
}
