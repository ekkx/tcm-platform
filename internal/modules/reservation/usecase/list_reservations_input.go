package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
	reservationv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/pkg/actor"
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
