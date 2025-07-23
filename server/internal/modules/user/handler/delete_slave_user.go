package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
)

func (h *HandlerImpl) DeleteSlaveUser(ctx context.Context, req *connect.Request[userv1.DeleteSlaveUserRequest]) (*connect.Response[userv1.DeleteSlaveUserResponse], error) {
    return nil, nil
}
