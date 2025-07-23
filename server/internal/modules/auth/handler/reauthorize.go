package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1"
)

func (h *HandlerImpl) Reauthorize(ctx context.Context, req *connect.Request[authv1.ReauthorizeRequest]) (*connect.Response[authv1.ReauthorizeResponse], error) {
    return nil, nil
}
