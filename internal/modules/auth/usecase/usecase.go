package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
)

type UseCase interface {
	Authorize(ctx context.Context, params *AuthorizeInput) (*AuthorizeOutput, error)
	Reauthorize(ctx context.Context, input *ReauthorizeInput) (*ReauthorizeOutput, error)
}

type UseCaseImpl struct {
	jwtManager *jwt.JWTManager
	userQuery  gateway.UserQuery
	userCmd    gateway.UserCommand
}

func New(
	jwtManager *jwt.JWTManager,
	userQuery gateway.UserQuery,
	userCmd gateway.UserCommand,
) UseCase {
	return &UseCaseImpl{
		jwtManager: jwtManager,
		userQuery:  userQuery,
		userCmd:    userCmd,
	}
}
