package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type UpdateUserByIDParams struct {
	UserID               ulid.ULID
	Password             *string
	OfficialSitePassword *string
	DisplayName          *string
}

func (repo *RepositoryImpl) UpdateUserByID(ctx context.Context, params *UpdateUserByIDParams) error {
	_, err := repo.querier.UpdateUserByID(ctx, database.UpdateUserByIDParams{
		UserID:               params.UserID,
		Password:             params.Password,
		OfficialSitePassword: params.OfficialSitePassword,
		DisplayName:          params.DisplayName,
	})
	if err != nil {
		return err
	}
	return nil
}
