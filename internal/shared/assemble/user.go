package assemble

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type UserView struct {
	User       entity.User
	MasterUser *entity.User
}

type UserAssembler interface {
	Build(ctx context.Context, userID ulid.ULID) (*UserView, error)
	BuildList(ctx context.Context, userIDs []ulid.ULID) ([]*UserView, error)
}

type UserAssemblerImpl struct {
	userQuery gateway.UserQuery
}

func NewUserAssembler(userQuery gateway.UserQuery) UserAssembler {
	return &UserAssemblerImpl{
		userQuery: userQuery,
	}
}

func (asm *UserAssemblerImpl) Build(ctx context.Context, userID ulid.ULID) (*UserView, error) {
	users, err := asm.BuildList(ctx, []ulid.ULID{userID})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}

func (asm *UserAssemblerImpl) BuildList(ctx context.Context, userIDs []ulid.ULID) ([]*UserView, error) {
	if len(userIDs) == 0 {
		return []*UserView{}, nil
	}

	entUsers, err := asm.userQuery.ListUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	// 重複を避けるためのセットを使用
	masterUserIDSet := make(map[ulid.ULID]struct{})
	for _, user := range entUsers {
		if user.MasterUserID != nil {
			masterUserIDSet[*user.MasterUserID] = struct{}{}
		}
	}

	masterUserIDs := make([]ulid.ULID, 0, len(masterUserIDSet))
	for id := range masterUserIDSet {
		masterUserIDs = append(masterUserIDs, id)
	}

	entMasterUsers, err := asm.userQuery.ListUsersByIDs(ctx, masterUserIDs)
	if err != nil {
		return nil, err
	}

	masterUserMap := make(map[ulid.ULID]*entity.User)
	for _, masterUser := range entMasterUsers {
		masterUserMap[masterUser.ID] = masterUser
	}

	out := make([]*UserView, 0, len(entUsers))
	for _, user := range entUsers {
		completeUser := &UserView{
			User:       *user,
			MasterUser: nil,
		}
		if user.MasterUserID != nil {
			if masterUser, exists := masterUserMap[*user.MasterUserID]; exists {
				completeUser.MasterUser = masterUser
			}
		}
		out = append(out, completeUser)
	}
	return out, nil
}
