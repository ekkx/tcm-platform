package usecase

import (
	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
)

type DeleteReservationOutput struct {
}

func NewDeleteReservationOutput() *DeleteReservationOutput {
	return &DeleteReservationOutput{}
}

func (st *DeleteReservationOutput) ToResponse() *connect.Response[reservationv1.DeleteReservationResponse] {
	return connect.NewResponse(&reservationv1.DeleteReservationResponse{})
}
