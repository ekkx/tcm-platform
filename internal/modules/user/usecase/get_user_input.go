package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/pkg/actor"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type GetUserInput struct {
	Actor  actor.Actor
	UserID ulid.ULID
}

func NewGetUserInputFromRequest(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (*GetUserInput, error) {
	st := &GetUserInput{}

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

func (st *GetUserInput) Validate() error {
	if st.UserID.IsZero() {
		return errs.ErrInvalidIDFormat
	}
	return nil
}
