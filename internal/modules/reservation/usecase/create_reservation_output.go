package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	reservationv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
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
