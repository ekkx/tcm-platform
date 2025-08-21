package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
)

type Repository interface {
	ListUnavailableRoomIDs(ctx context.Context, params *ListUnavailableRoomIDsParams) ([]string, error)
}

type RepositoryImpl struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
