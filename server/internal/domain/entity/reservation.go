package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
)

type Reservation struct {
	ID         int             `json:"id"`
	ExternalID *string         `json:"external_id"`
	UserID     string          `json:"user_id"`
	CampusType enum.CampusType `json:"campus_type"`
	RoomID     string          `json:"room_id"`
	Date       time.Time       `json:"date"`
	FromHour   int             `json:"from_hour"`
	FromMinute int             `json:"from_minute"`
	ToHour     int             `json:"to_hour"`
	ToMinute   int             `json:"to_minute"`
	BookerName *string         `json:"booker_name"`
	CreatedAt  time.Time       `json:"created_at"`
}
