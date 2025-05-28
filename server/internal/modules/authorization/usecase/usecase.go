package usecase

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/port"
	userRepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
)

type Usecase struct {
	tcmClient port.TCMClient
	userRepo  *userRepo.Repository
}

func NewUsecase(tcmClient port.TCMClient, userRepo *userRepo.Repository) *Usecase {
	return &Usecase{
		tcmClient: tcmClient,
		userRepo:  userRepo,
	}
}
