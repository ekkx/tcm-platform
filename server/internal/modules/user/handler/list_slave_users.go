package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/usecase"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
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
