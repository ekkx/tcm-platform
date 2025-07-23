package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/usecase"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
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
