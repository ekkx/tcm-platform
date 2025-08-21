package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
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
