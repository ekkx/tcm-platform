package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
)

func (h *Handler) CreateSlaveUser(ctx context.Context, req *connect.Request[userv1.CreateSlaveUserRequest]) (*connect.Response[userv1.CreateSlaveUserResponse], error) {
    return nil, nil
}
