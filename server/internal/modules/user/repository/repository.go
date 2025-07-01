package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error)
}

type RepositoryImpl struct {
	querier database.Querier
}

func New(querier database.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
