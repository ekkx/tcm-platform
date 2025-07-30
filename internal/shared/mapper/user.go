package mapper

import (
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/database"
)

func ToUser(user *database.User) *entity.User {
	if user == nil {
		return nil
	}

	var masterUser *entity.User
	if user.MasterUserID != nil {
		masterUser = &entity.User{ID: *user.MasterUserID}
	}

	return &entity.User{
		ID:                   user.ID,
		Password:             user.Password,
		OfficialSiteID:       user.OfficialSiteID,
		OfficialSitePassword: user.OfficialSitePassword,
		MasterUser:           masterUser,
		DisplayName:          user.DisplayName,
		CreateTime:           user.CreateTime,
	}
}
