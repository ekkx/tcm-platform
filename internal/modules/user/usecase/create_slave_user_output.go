package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
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
