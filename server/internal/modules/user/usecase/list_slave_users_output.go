package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type ListSlaveUsersOutput struct {
	Users []*entity.User
}

func NewListSlaveUsersOutput(users []*entity.User) *ListSlaveUsersOutput {
	return &ListSlaveUsersOutput{
		Users: users,
	}
}

func (st *ListSlaveUsersOutput) ToResponse() *connect.Response[userv1.ListSlaveUsersResponse] {
	return connect.NewResponse(&userv1.ListSlaveUsersResponse{
		Users: presenter.ToUserList(st.Users),
	})
}
