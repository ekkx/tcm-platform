package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/enum"
)

type Reservation struct {
	ID         int64           `json:"id"`
	ExternalID *string         `json:"external_id"`
	UserID     string          `json:"user_id"`
	CampusType enum.CampusType `json:"campus_type"`
	RoomID     string          `json:"room_id"`
	Date       time.Time       `json:"date"`
	FromHour   int32           `json:"from_hour"`
	FromMinute int32           `json:"from_minute"`
	ToHour     int32           `json:"to_hour"`
	ToMinute   int32           `json:"to_minute"`
	BookerName *string         `json:"booker_name"`
	CreatedAt  time.Time       `json:"created_at"`
}
