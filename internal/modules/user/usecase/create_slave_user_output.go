package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
)

type CreateSlaveUserOutput struct {
	User assemble.UserView
}

func NewCreateSlaveUserOutput(user assemble.UserView) *CreateSlaveUserOutput {
	return &CreateSlaveUserOutput{
		User: user,
	}
}

func (st *CreateSlaveUserOutput) ToResponse() *connect.Response[userv1.CreateSlaveUserResponse] {
	return connect.NewResponse(&userv1.CreateSlaveUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
