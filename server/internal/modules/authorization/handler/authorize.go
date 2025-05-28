package handler

import (
	"context"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
)

func (h *Handler) Authorize(
	ctx context.Context,
	req *auth_v1.AuthorizeRequest,
) (*auth_v1.AuthorizeReply, error) {
	output, err := h.Usecase.Authorize(ctx, input.NewAuthorize().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
