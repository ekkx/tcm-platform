package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/pkg/actor"
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
