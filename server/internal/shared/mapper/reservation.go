package mapper

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
)

func ToReservation(rsv *database.Reservation) *entity.Reservation {
	var campusType enum.CampusType
	switch rsv.CampusType {
	case database.CampusTypeIkebukuro:
		campusType = enum.CampusTypeIkebukuro
	case database.CampusTypeNakameguro:
		campusType = enum.CampusTypeNakameguro
	default:
		campusType = enum.CampusTypeUnknown
	}

	return &entity.Reservation{
		ID:             rsv.ID,
		OfficialSiteID: rsv.OfficialSiteID,
		User:           entity.User{ID: rsv.UserID},
		CampusType:     campusType,
		Room:           entity.Room{ID: rsv.RoomID},
		Date:           rsv.Date,
		FromHour:       int(rsv.FromHour),
		FromMinute:     int(rsv.FromMinute),
		ToHour:         int(rsv.ToHour),
		ToMinute:       int(rsv.ToMinute),
		CreateTime:     rsv.CreateTime,
	}
}
