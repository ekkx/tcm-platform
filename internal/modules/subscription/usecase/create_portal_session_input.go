package usecase

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type CreatePortalSessionInput struct {
	Actor actor.Actor
}

func NewCreatePortalSessionInputFromRequest(ctx context.Context, _ *connect.Request[subscriptionv1.CreatePortalSessionRequest]) (*CreatePortalSessionInput, error) {
	a := ctxhelper.Actor(ctx)
	if a == nil {
		return nil, errs.ErrUnauthorized
	}
	return &CreatePortalSessionInput{Actor: *a}, nil
}
