package usecase

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/platform/actor"
	"github.com/ekkx/tcm-platform/internal/platform/ctxhelper"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

type CreateCheckoutSessionInput struct {
	Actor   actor.Actor
	PriceID string
}

func NewCreateCheckoutSessionInputFromRequest(ctx context.Context, req *connect.Request[subscriptionv1.CreateCheckoutSessionRequest]) (*CreateCheckoutSessionInput, error) {
	a := ctxhelper.Actor(ctx)
	if a == nil {
		return nil, errs.ErrUnauthorized
	}

	priceID := req.Msg.PriceId
	if priceID == "" {
		return nil, errs.ErrInvalidArgument.WithMessage("price_id is required")
	}

	return &CreateCheckoutSessionInput{
		Actor:   *a,
		PriceID: priceID,
	}, nil
}
