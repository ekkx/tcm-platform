package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
	"github.com/ekkx/tcmrsv-web/pkg/ymd"
)

type Reservation struct {
	ID             ulid.ULID
	OfficialSiteID *string
	User           User
	CampusType     enum.CampusType
	Room           Room
	Date           ymd.YMD
	FromHour       int
	FromMinute     int
	ToHour         int
	ToMinute       int
	CreateTime     time.Time
}
