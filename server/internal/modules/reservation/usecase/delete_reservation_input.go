package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/pkg/actor"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type DeleteReservationInput struct {
	Actor         actor.Actor
	ReservationID ulid.ULID
}

func NewDeleteReservationInputFromRequest(ctx context.Context, req *connect.Request[reservationv1.DeleteReservationRequest]) (*DeleteReservationInput, error) {
	st := &DeleteReservationInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	parsedID, err := ulid.Parse(req.Msg.ReservationId)
	if err != nil {
		parsedID = ulid.ULID{}
	}

	st.ReservationID = parsedID

	return st, nil
}

func (st *DeleteReservationInput) Validate() error {
	if st.ReservationID.IsZero() {
		return errs.ErrInvalidIDFormat
	}
	return nil
}
