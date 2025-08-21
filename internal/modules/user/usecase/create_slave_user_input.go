package usecase

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type CreateSlaveUserInput struct {
	Actor       actor.Actor
	DisplayName string
	Password    string
}

func NewCreateSlaveUserInputFromRequest(ctx context.Context, req *connect.Request[userv1.CreateSlaveUserRequest]) (*CreateSlaveUserInput, error) {
	st := &CreateSlaveUserInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	st.DisplayName = req.Msg.DisplayName
	st.Password = req.Msg.Password

	return st, nil
}

func (st *CreateSlaveUserInput) Validate() error {
	return nil
}
