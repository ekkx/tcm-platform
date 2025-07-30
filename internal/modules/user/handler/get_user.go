package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/usecase"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
)

func (h *HandlerImpl) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error) {
	input, err := usecase.NewGetUserInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.GetUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
