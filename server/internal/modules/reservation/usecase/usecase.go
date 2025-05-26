package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"

type Usecase struct {
	rsvrepo *repository.Repository
}

func NewUsecase(
	rsvrepo *repository.Repository,
) *Usecase {
	return &Usecase{
		rsvrepo: rsvrepo,
	}
}
