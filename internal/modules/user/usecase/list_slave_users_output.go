package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
)

type ListSlaveUsersOutput struct {
	Users []*assemble.UserView
}

func NewListSlaveUsersOutput(users []*assemble.UserView) *ListSlaveUsersOutput {
	return &ListSlaveUsersOutput{
		Users: users,
	}
}

func (st *ListSlaveUsersOutput) ToResponse() *connect.Response[userv1.ListSlaveUsersResponse] {
	return connect.NewResponse(&userv1.ListSlaveUsersResponse{
		Users: presenter.ToUserList(st.Users),
	})
}
