package usecase

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type UpdateReservationNoteInput struct {
	Actor         actor.Actor
	ReservationID ulid.ULID
	Note          *string
}

func NewUpdateReservationNoteInputFromRequest(ctx context.Context, req *connect.Request[reservationv1.UpdateReservationNoteRequest]) (*UpdateReservationNoteInput, error) {
	st := &UpdateReservationNoteInput{}

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
	st.Note = req.Msg.Note

	return st, nil
}

func (st *UpdateReservationNoteInput) Validate() error {
	if st.ReservationID.IsZero() {
		return errs.ErrInvalidIDFormat
	}
	return nil
}
