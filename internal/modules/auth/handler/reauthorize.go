package handler

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
	"github.com/ekkx/tcm-platform/internal/modules/auth/usecase"
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
