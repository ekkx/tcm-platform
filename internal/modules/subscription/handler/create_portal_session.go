package handler

import (
	"context"

	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/usecase"
)

func (h *HandlerImpl) CreatePortalSession(ctx context.Context, req *connect.Request[subscriptionv1.CreatePortalSessionRequest]) (*connect.Response[subscriptionv1.CreatePortalSessionResponse], error) {
	input, err := usecase.NewCreatePortalSessionInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.CreatePortalSession(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
