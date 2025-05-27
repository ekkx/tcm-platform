package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	auth_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorization"
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
