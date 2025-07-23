package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1/authv1connect"
)

type HandlerImpl struct {
	authUseCase usecase.UseCase
}

func New(authUseCase usecase.UseCase) authv1connect.AuthServiceHandler {
	return &HandlerImpl{
		authUseCase: authUseCase,
	}
}
