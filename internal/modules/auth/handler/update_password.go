package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
)

func (h *HandlerImpl) UpdatePassword(ctx context.Context, req *connect.Request[authv1.UpdatePasswordRequest]) (*connect.Response[authv1.UpdatePasswordResponse], error) {
	return nil, nil
}
