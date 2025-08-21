package usecase

import (
	"connectrpc.com/connect"
	userv1 "github.com/ekkx/tcm-platform/internal/gen/pb/user/v1"
)

type DeleteUserOutput struct {
}

func NewDeleteUserOutput() *DeleteUserOutput {
	return &DeleteUserOutput{}
}

func (st *DeleteUserOutput) ToResponse() *connect.Response[userv1.DeleteUserResponse] {
	return connect.NewResponse(&userv1.DeleteUserResponse{})
}
