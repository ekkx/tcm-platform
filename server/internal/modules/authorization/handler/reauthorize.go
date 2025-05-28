package handler

import (
	"context"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
)

func (h *Handler) Reauthorize(
	ctx context.Context,
	req *auth_v1.ReauthorizeRequest,
) (*auth_v1.ReauthorizeReply, error) {
	output, err := h.Usecase.Reauthorize(ctx, input.NewReauthorize().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
