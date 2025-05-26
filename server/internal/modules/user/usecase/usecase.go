package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"

type Usecase struct {
	userrepo *repository.Repository
}

func NewUsecase(
	userrepo *repository.Repository,
) *Usecase {
	return &Usecase{
		userrepo: userrepo,
	}
}
