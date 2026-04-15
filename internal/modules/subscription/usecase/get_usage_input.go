package usecase

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type GetUsageInput struct {
	Actor actor.Actor
}

func NewGetUsageInputFromRequest(ctx context.Context, _ *connect.Request[subscriptionv1.GetUsageRequest]) (*GetUsageInput, error) {
	a := ctxhelper.Actor(ctx)
	if a == nil {
		return nil, errs.ErrUnauthorized
	}
	return &GetUsageInput{Actor: *a}, nil
}
