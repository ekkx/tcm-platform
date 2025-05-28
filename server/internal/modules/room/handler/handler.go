package handler

import (
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/room"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
)

type Handler struct {
	room_v1.UnimplementedRoomServiceServer

	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
