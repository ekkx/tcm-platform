package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
)

func (h *HandlerImpl) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[userv1.DeleteUserResponse], error) {
	input, err := usecase.NewDeleteUserInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.DeleteUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
