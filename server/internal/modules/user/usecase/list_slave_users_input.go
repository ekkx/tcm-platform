package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/server/pkg/actor"
)

type ListSlaveUsersInput struct {
	Actor actor.Actor
}

func NewListSlaveUsersInputFromRequest(ctx context.Context, req *connect.Request[userv1.ListSlaveUsersRequest]) (*ListSlaveUsersInput, error) {
	st := &ListSlaveUsersInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	return st, nil
}

func (st *ListSlaveUsersInput) Validate() error {
	return nil
}
