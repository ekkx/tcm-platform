package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
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
