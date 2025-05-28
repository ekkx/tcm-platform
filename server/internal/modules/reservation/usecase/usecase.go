package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"

type Usecase struct {
	rsvRepo *repository.Repository
}

func NewUsecase(
	rsvRepo *repository.Repository,
) *Usecase {
	return &Usecase{
		rsvRepo: rsvRepo,
	}
}
