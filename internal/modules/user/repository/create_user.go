package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type CreateUserParams struct {
	ID                   ulid.ULID
	Password             string
	OfficialSiteID       *string
	OfficialSitePassword *string
	MasterUserID         *ulid.ULID
	DisplayName          string
}

func (repo *RepositoryImpl) CreateUser(ctx context.Context, params *CreateUserParams) (*ulid.ULID, error) {
	if params.ID.IsZero() {
		params.ID = ulid.New()
	}

	if params.DisplayName == "" {
		params.DisplayName = "未設定" // デフォルトの表示名を設定
	}

	id, err := repo.querier.CreateUser(ctx, sqlc.CreateUserParams{
		ID:                   params.ID,
		Password:             params.Password,
		OfficialSiteID:       params.OfficialSiteID,
		OfficialSitePassword: params.OfficialSitePassword,
		MasterUserID:         params.MasterUserID,
		DisplayName:          params.DisplayName,
	})
	if err != nil {
		return nil, err
	}

	return &id, nil
}
