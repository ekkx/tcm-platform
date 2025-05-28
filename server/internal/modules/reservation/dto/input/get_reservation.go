package input

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type GetReservation struct {
	Actor         actor.Actor
	ReservationID int64
}

func NewGetReservation() *GetReservation {
	return &GetReservation{}
}

func (input *GetReservation) FromProto(ctx context.Context, req *rsv_v1.GetReservationRequest) *GetReservation {
	input.Actor = ctxhelper.GetActor(ctx)
	input.ReservationID = req.ReservationId
	return input
}
