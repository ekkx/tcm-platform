package usecase

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type DeleteUserInput struct {
	Actor  actor.Actor
	UserID ulid.ULID
}

func NewDeleteUserInputFromRequest(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*DeleteUserInput, error) {
	st := &DeleteUserInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	parsedID, err := ulid.Parse(req.Msg.UserId)
	if err != nil {
		parsedID = ulid.ULID{}
	}

	st.UserID = parsedID

	return st, nil
}

func (st *DeleteUserInput) Validate() error {
	if st.UserID.IsZero() {
		return errs.ErrInvalidIDFormat
	}
	return nil
}
