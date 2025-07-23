package handler

import (
	"context"

	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
)

func (h *HandlerImpl) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[userv1.DeleteUserResponse], error) {
    return nil, nil
}
