package entity

import (
	"time"

	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type Reservation struct {
	ID             ulid.ULID
	OfficialSiteID *string
	UserID         ulid.ULID
	CampusType     valueobject.CampusType
	RoomID         string
	Date           ymd.YMD
	FromHour       int
	FromMinute     int
	ToHour         int
	ToMinute       int
	Status         valueobject.ReservationStatus
	CreateTime     time.Time
}
