package domain

import "time"

type Reservation struct {
	ID         string    `json:"id"`
	RoomID     string    `json:"room_id"`
	Date       time.Time `json:"date"`
	FromHour   int       `json:"from_hour"`
	FromMinute int       `json:"from_minute"`
	ToHour     int       `json:"to_hour"`
	ToMinute   int       `json:"to_minute"`
	BookerName *string   `json:"booker_name"`
	CreatedAt  time.Time `json:"created_at"`
}
