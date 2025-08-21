package handler

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/auth/usecase"
)

type HandlerImpl struct {
	authUseCase usecase.UseCase
}

func New(authUseCase usecase.UseCase) authv1connect.AuthServiceHandler {
	return &HandlerImpl{
		authUseCase: authUseCase,
	}
}
