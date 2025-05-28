package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"

type Usecase struct {
	userRepo *repository.Repository
}

func NewUsecase(
	userRepo *repository.Repository,
) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}
