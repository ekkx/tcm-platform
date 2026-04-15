package handler

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/usecase"
)

func (h *HandlerImpl) GetSubscription(ctx context.Context, req *connect.Request[subscriptionv1.GetSubscriptionRequest]) (*connect.Response[subscriptionv1.GetSubscriptionResponse], error) {
	input, err := usecase.NewGetSubscriptionInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.GetSubscription(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
