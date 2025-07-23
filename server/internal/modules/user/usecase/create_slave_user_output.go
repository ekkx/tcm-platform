package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type CreateSlaveUserOutput struct {
	entity.User
}

func NewCreateSlaveUserOutput(user entity.User) *CreateSlaveUserOutput {
	return &CreateSlaveUserOutput{
		User: user,
	}
}

func (st *CreateSlaveUserOutput) ToResponse() *connect.Response[userv1.CreateSlaveUserResponse] {
	return connect.NewResponse(&userv1.CreateSlaveUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
