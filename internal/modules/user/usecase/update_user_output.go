package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type UpdateUserOutput struct {
	User assemble.UserView
}

func NewUpdateUserOutput(user assemble.UserView) *UpdateUserOutput {
	return &UpdateUserOutput{
		User: user,
	}
}

func (st *UpdateUserOutput) ToResponse() *connect.Response[userv1.UpdateUserResponse] {
	return connect.NewResponse(&userv1.UpdateUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
