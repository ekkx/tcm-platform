package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type AuthorizeOutput struct {
	entity.Auth
}

func NewAuthorizeOutput(auth entity.Auth) *AuthorizeOutput {
	return &AuthorizeOutput{
		Auth: auth,
	}
}

func (st *AuthorizeOutput) ToResponse() *connect.Response[authv1.AuthorizeResponse] {
	return connect.NewResponse(&authv1.AuthorizeResponse{
		Auth: presenter.ToAuth(&st.Auth),
	})
}
