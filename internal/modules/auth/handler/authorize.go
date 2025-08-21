package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
	"github.com/ekkx/tcm-platform/internal/modules/auth/usecase"
)

func (h *HandlerImpl) Authorize(ctx context.Context, req *connect.Request[authv1.AuthorizeRequest]) (*connect.Response[authv1.AuthorizeResponse], error) {
	input, err := usecase.NewAuthorizeInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.authUseCase.Authorize(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
