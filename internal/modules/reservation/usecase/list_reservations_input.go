package usecase

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type ListReservationsInput struct {
	Actor actor.Actor
}

func NewListReservationsInputFromRequest(ctx context.Context, req *connect.Request[reservationv1.ListReservationsRequest]) (*ListReservationsInput, error) {
	st := &ListReservationsInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	return st, nil
}

func (st *ListReservationsInput) Validate() error {
	return nil
}
