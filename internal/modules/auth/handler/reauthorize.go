package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/auth/usecase"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
)

func (h *HandlerImpl) Reauthorize(ctx context.Context, req *connect.Request[authv1.ReauthorizeRequest]) (*connect.Response[authv1.ReauthorizeResponse], error) {
	input, err := usecase.NewReauthorizeInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.authUseCase.Reauthorize(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
