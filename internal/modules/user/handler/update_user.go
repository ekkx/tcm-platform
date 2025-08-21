package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
)

func (h *HandlerImpl) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error) {
	input, err := usecase.NewUpdateUserInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.UpdateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
