package repository

import (
	"context"
	"errors"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error) {
	// ユーザーの骨組みとなる情報を取得
	userMeta, err := repo.querier.GetUserMetaByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	ids := []ulid.ULID{userMeta.UserID}
	if userMeta.MasterUserID != nil {
		ids = append(ids, *userMeta.MasterUserID)
	}

	// ユーザーの骨組み情報をもとに、肉付け用の詳細情報を取得
	users, err := repo.querier.ListUsersByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	// TODO: mapper などに切り出す
	userMap := make(map[ulid.ULID]*entity.User)
	for _, u := range users {
		userMap[u.ID] = &entity.User{
			ID:          u.ID,
			DisplayName: u.DisplayName,
			MasterUser:  nil, // のちにマスターユーザーが設定される
			CreateTime:  u.CreateTime,
		}
	}

	user := userMap[userMeta.UserID]
	if user == nil {
		return nil, nil
	}

	if userMeta.MasterUserID != nil {
		if masterUser, ok := userMap[*userMeta.MasterUserID]; ok {
			user.MasterUser = masterUser
		}
	}

	return user, nil
}
