package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type UpdateUserByIDParams struct {
	ID                   ulid.ULID
	Password             *string
	OfficialSitePassword *string
	DisplayName          *string
}

func (repo *RepositoryImpl) UpdateUserByID(ctx context.Context, params *UpdateUserByIDParams) error {
	_, err := repo.querier.UpdateUserByID(ctx, database.UpdateUserByIDParams{
		UserID:               params.ID,
		Password:             params.Password,
		OfficialSitePassword: params.OfficialSitePassword,
		DisplayName:          params.DisplayName,
	})
	if err != nil {
		return err
	}
	return nil
}
