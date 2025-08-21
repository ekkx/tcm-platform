package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
)

type GetUserOutput struct {
	User assemble.UserView
}

func NewGetUserOutput(user assemble.UserView) *GetUserOutput {
	return &GetUserOutput{
		User: user,
	}
}

func (st *GetUserOutput) ToResponse() *connect.Response[userv1.GetUserResponse] {
	return connect.NewResponse(&userv1.GetUserResponse{
		User: presenter.ToUser(&st.User),
	})
}
