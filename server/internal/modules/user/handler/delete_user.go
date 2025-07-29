package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/usecase"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
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
