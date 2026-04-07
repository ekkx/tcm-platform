package valueobject

type ReservationStatus int

const (
	ReservationStatusPending ReservationStatus = iota
	ReservationStatusSuccess
	ReservationStatusFailed
)
