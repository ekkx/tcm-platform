package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type ReauthorizeOutput struct {
	entity.Auth
}

func NewReauthorizeOutput(auth entity.Auth) *ReauthorizeOutput {
	return &ReauthorizeOutput{
		Auth: auth,
	}
}

func (st *ReauthorizeOutput) ToResponse() *connect.Response[authv1.ReauthorizeResponse] {
	return connect.NewResponse(&authv1.ReauthorizeResponse{
		Auth: presenter.ToAuth(&st.Auth),
	})
}
