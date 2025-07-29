package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type GetReservationOutput struct {
	Reservation entity.Reservation
}

func NewGetReservationOutput(reservation entity.Reservation) *GetReservationOutput {
	return &GetReservationOutput{
		Reservation: reservation,
	}
}

func (st *GetReservationOutput) ToResponse() *connect.Response[reservationv1.GetReservationResponse] {
	return connect.NewResponse(&reservationv1.GetReservationResponse{
		Reservation: presenter.ToReservation(&st.Reservation),
	})
}
