package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
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
