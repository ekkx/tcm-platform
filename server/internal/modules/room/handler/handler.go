package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"
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
