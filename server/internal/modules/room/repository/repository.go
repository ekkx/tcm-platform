package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/pkg/database"
)

type Repository interface {
	ListUnavailableRoomIDs(ctx context.Context, params *ListUnavailableRoomIDsParams) ([]string, error)
}

type RepositoryImpl struct {
	querier database.Querier
}

func New(querier database.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
