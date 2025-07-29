package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/service"
)

type UseCase interface {
	CreateSlaveUser(ctx context.Context, params *CreateSlaveUserInput) (*CreateSlaveUserOutput, error)
	DeleteUser(ctx context.Context, input *DeleteUserInput) (*DeleteUserOutput, error)
}

type UseCaseImpl struct {
	userRepo    repository.Repository
	userService service.Service
}

func New(
	userRepo repository.Repository,
	userService service.Service,
) UseCase {
	return &UseCaseImpl{
		userRepo:    userRepo,
		userService: userService,
	}
}
