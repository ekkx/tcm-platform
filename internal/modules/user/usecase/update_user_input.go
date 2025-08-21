package usecase

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type UpdateUserInput struct {
	Actor       actor.Actor
	DisplayName string
}

func NewUpdateUserInputFromRequest(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*UpdateUserInput, error) {
	st := &UpdateUserInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	st.DisplayName = req.Msg.DisplayName

	return st, nil
}

func (st *UpdateUserInput) Validate() error {
	return nil
}
