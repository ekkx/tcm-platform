package usecase

import (
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

type Usecase struct {
	rsvRepo  *rsv_repo.Repository
	roomRepo *room_repo.Repository
}

func NewUsecase(
	rsvRepo *rsv_repo.Repository,
	roomRepo *room_repo.Repository,
) *Usecase {
	return &Usecase{
		rsvRepo:  rsvRepo,
		roomRepo: roomRepo,
	}
}
