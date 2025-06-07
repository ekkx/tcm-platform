package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"

type Usecase interface {
	// TODO: Add user-related methods when implementing user handler
}

type UsecaseImpl struct {
	userRepo *repository.Repository
}

func NewUsecase(
	userRepo *repository.Repository,
) *UsecaseImpl {
	return &UsecaseImpl{
		userRepo: userRepo,
	}
}
