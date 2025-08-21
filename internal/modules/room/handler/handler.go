package handler

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/room/v1/roomv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/room/usecase"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) roomv1connect.RoomServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
