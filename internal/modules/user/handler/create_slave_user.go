package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
)

func (h *HandlerImpl) CreateSlaveUser(ctx context.Context, req *connect.Request[userv1.CreateSlaveUserRequest]) (*connect.Response[userv1.CreateSlaveUserResponse], error) {
	input, err := usecase.NewCreateSlaveUserInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.CreateSlaveUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
