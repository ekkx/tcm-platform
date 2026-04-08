package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type UpdateReservationNoteParams struct {
	ID   ulid.ULID
	Note *string
}

func (repo *RepositoryImpl) UpdateReservationNote(ctx context.Context, params *UpdateReservationNoteParams) error {
	return repo.querier.UpdateReservationNote(ctx, sqlc.UpdateReservationNoteParams{
		ID:   params.ID,
		Note: params.Note,
	})
}
