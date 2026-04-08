package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
)

type UpdateReservationNoteOutput struct {
	Reservation assemble.ReservationView
}

func NewUpdateReservationNoteOutput(v assemble.ReservationView) *UpdateReservationNoteOutput {
	return &UpdateReservationNoteOutput{
		Reservation: v,
	}
}

func (st *UpdateReservationNoteOutput) ToResponse() *connect.Response[reservationv1.UpdateReservationNoteResponse] {
	return connect.NewResponse(&reservationv1.UpdateReservationNoteResponse{
		Reservation: presenter.ToReservation(&st.Reservation),
	})
}
