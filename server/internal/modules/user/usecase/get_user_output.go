package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type GetUserOutput struct {
	User entity.User
}

func NewGetUserOutput(user entity.User) *GetUserOutput {
	return &GetUserOutput{
		User: user,
	}
}

func (st *GetUserOutput) ToResponse() *connect.Response[userv1.GetUserResponse] {
	return connect.NewResponse(&userv1.GetUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
