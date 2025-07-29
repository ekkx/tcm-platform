package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) DeleteUserByID(ctx context.Context, userID ulid.ULID) error {
	if _, err := repo.querier.DeleteUserByID(ctx, userID); err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrUserNotFound
		}
		return err
	}
	return nil
}
