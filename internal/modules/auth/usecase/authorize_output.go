package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
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
