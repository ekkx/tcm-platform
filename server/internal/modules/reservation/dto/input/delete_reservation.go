package input

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type DeleteReservation struct {
	Actor         actor.Actor
	ReservationID int
}

func NewDeleteReservation() *DeleteReservation {
	return &DeleteReservation{}
}

func (input *DeleteReservation) FromProto(ctx context.Context, req *reservation.DeleteReservationRequest) *DeleteReservation {
	input.Actor = ctxhelper.GetActor(ctx)
	input.ReservationID = int(req.ReservationId)
	return input
}
