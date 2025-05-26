package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorize/usecase"
	authorize_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorize"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
)

func (h *Handler) Authorize(
	ctx context.Context,
	req *authorize_v1.AuthorizeRequest,
) (*authorize_v1.AuthorizeReply, error) {
	output, err := h.Usecase.Authorize(ctx, &usecase.AuthorizeInput{
		UserID:         req.UserId,
		Password:       req.Password,
		PasswordAESKey: ctxhelper.GetConfig(ctx).PasswordAESKey,
		JWTSecret:      ctxhelper.GetConfig(ctx).JWTSecret,
	})
	if err != nil {
		return nil, err
	}

	return &authorize_v1.AuthorizeReply{
		Authorization: &authorize_v1.Authorization{
			AccessToken:  output.Authorization.AccessToken,
			RefreshToken: output.Authorization.RefreshToken,
		},
	}, nil
}
