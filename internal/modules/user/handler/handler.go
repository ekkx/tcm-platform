package handler

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/user/v1/userv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) userv1connect.UserServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
