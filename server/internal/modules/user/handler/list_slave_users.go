package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
)

func (h *Handler) ListSlaveUsers(ctx context.Context, req *connect.Request[userv1.ListSlaveUsersRequest]) (*connect.Response[userv1.ListSlaveUsersResponse], error) {
    return nil, nil
}
