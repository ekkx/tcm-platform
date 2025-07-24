package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type Reservation struct {
	ID             int
	OfficialSiteID *string
	User           User
	CampusType     enum.CampusType
	RoomID         string
	Date           ymd.YMD
	FromHour       int
	FromMinute     int
	ToHour         int
	ToMinute       int
	CreateTime     time.Time
}
