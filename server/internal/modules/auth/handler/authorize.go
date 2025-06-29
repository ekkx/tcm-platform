package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1"
)

func (h *Handler) Authorize(ctx context.Context, req *connect.Request[authv1.AuthorizeRequest]) (*connect.Response[authv1.AuthorizeResponse], error) {
    return nil, nil
}
