package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
)

func (h *HandlerImpl) ListSlaveUsers(ctx context.Context, req *connect.Request[userv1.ListSlaveUsersRequest]) (*connect.Response[userv1.ListSlaveUsersResponse], error) {
	input, err := usecase.NewListSlaveUsersInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.ListSlaveUsers(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
