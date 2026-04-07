package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type UpdateReservationStatusParams struct {
	ID             ulid.ULID
	Status         valueobject.ReservationStatus
	OfficialSiteID *string
}

func (repo *RepositoryImpl) UpdateReservationStatus(ctx context.Context, params *UpdateReservationStatusParams) error {
	var status sqlc.ReservationStatus
	switch params.Status {
	case valueobject.ReservationStatusSuccess:
		status = sqlc.ReservationStatusSuccess
	case valueobject.ReservationStatusFailed:
		status = sqlc.ReservationStatusFailed
	default:
		status = sqlc.ReservationStatusPending
	}

	return repo.querier.UpdateReservationStatus(ctx, sqlc.UpdateReservationStatusParams{
		ID:             params.ID,
		Status:         status,
		OfficialSiteID: params.OfficialSiteID,
	})
}
