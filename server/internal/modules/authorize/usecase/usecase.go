package usecase

import (
	"github.com/ekkx/tcmrsv"
	userRepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
)

type Usecase struct {
	userRepo *userRepo.Repository

	tcmClient *tcmrsv.Client
}

func NewUsecase(userrepo *userRepo.Repository, tcmClient *tcmrsv.Client) *Usecase {
	return &Usecase{
		userRepo:  userrepo,
		tcmClient: tcmClient,
	}
}
