package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) DeleteReservationByID(ctx context.Context, reservationID ulid.ULID) error {
	if _, err := repo.querier.DeleteReservationByID(ctx, reservationID); err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrReservationNotFound
		}
		return err
	}
	return nil
}
