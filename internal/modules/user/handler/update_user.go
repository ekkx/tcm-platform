package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/usecase"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
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
