package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"

type Usecase interface{}

type UsecaseImpl struct {
	userRepository repository.Repository
}

func New(userRepository repository.Repository) Usecase {
	return &UsecaseImpl{
		userRepository: userRepository,
	}
}
