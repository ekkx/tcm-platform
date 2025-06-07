package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

type Usecase interface {
	GetRooms(ctx context.Context) *output.GetRooms
}

type UsecaseImpl struct {
	roomRepo *repository.Repository
}

func NewUsecase(roomRepo *repository.Repository) *UsecaseImpl {
	return &UsecaseImpl{
		roomRepo: roomRepo,
	}
}
