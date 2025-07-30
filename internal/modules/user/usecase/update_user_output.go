package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type UpdateUserOutput struct {
	User entity.User
}

func NewUpdateUserOutput(user entity.User) *UpdateUserOutput {
	return &UpdateUserOutput{
		User: user,
	}
}

func (st *UpdateUserOutput) ToResponse() *connect.Response[userv1.UpdateUserResponse] {
	return connect.NewResponse(&userv1.UpdateUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
