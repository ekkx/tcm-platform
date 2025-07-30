package usecase

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
)

type AuthorizeInput struct {
	UserID   string // 公式サイトのユーザーIDも使われる想定なので ulid.ULID ではない
	Password string
}

func NewAuthorizeInputFromRequest(ctx context.Context, req *connect.Request[authv1.AuthorizeRequest]) (*AuthorizeInput, error) {
	return &AuthorizeInput{
		UserID:   req.Msg.UserId,
		Password: req.Msg.Password,
	}, nil
}

func (st *AuthorizeInput) Validate() error {
	return nil
}
