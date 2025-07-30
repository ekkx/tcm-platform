package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/auth/usecase"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
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
