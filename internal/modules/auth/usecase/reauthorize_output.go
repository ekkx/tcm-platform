package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
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
