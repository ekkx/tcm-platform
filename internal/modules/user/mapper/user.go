package mapper

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
)

func ToUser(user *sqlc.User) *entity.User {
	if user == nil {
		return nil
	}

	return &entity.User{
		ID:                   user.ID,
		Password:             user.Password,
		OfficialSiteID:       user.OfficialSiteID,
		OfficialSitePassword: user.OfficialSitePassword,
		MasterUserID:         user.MasterUserID,
		DisplayName:          user.DisplayName,
		CreateTime:           user.CreateTime,
	}
}
