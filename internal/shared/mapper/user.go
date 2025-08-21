package mapper

import (
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/database"
)

func ToUser(user *database.User) *entity.User {
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
