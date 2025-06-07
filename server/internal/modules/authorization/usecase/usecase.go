package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/port"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/output"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
)

type Usecase interface {
	Authorize(ctx context.Context, params *input.Authorize) (*output.Authorize, error)
    Reauthorize(ctx context.Context, params *input.Reauthorize) (*output.Reauthorize, error)
}

type UsecaseImpl struct {
	tcmClient port.TCMClient
	userRepo  *user_repo.Repository
}

func NewUsecase(tcmClient port.TCMClient, userRepo *user_repo.Repository) *UsecaseImpl {
	return &UsecaseImpl{
		tcmClient: tcmClient,
		userRepo:  userRepo,
	}
}
