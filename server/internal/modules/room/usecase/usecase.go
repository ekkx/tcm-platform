package usecase

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

type Usecase struct {
	roomRepo *repository.Repository
}

func NewUsecase(roomRepo *repository.Repository) *Usecase {
	return &Usecase{
		roomRepo: roomRepo,
	}
}
