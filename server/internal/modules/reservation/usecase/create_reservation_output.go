package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type CreateReservationOutput struct {
	Reservation entity.Reservation
}

func NewCreateReservationOutput(reservation entity.Reservation) *CreateReservationOutput {
	return &CreateReservationOutput{
		Reservation: reservation,
	}
}

func (st *CreateReservationOutput) ToResponse() *connect.Response[reservationv1.CreateReservationResponse] {
	return connect.NewResponse(&reservationv1.CreateReservationResponse{
		Reservation: presenter.ToReservation(&st.Reservation),
	})
}
