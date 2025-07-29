package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1/roomv1connect"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) roomv1connect.RoomServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
