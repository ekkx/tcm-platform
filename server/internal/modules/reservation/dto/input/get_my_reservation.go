package input

import "time"

type GetMyReservations struct {
	UserID string
	Date   time.Time
}
