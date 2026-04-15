package handler

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/usecase"
)

func (h *HandlerImpl) CreateCheckoutSession(ctx context.Context, req *connect.Request[subscriptionv1.CreateCheckoutSessionRequest]) (*connect.Response[subscriptionv1.CreateCheckoutSessionResponse], error) {
	input, err := usecase.NewCreateCheckoutSessionInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.CreateCheckoutSession(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
