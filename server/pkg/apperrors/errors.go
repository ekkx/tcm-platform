package apperrors

import "errors"

var (
	// auth
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")

	// users
	ErrUserNotFound        = errors.New("user not found")
	ErrRequestUserNotFound = errors.New("request user not found")

	// reservations
	ErrReservationNotFound = errors.New("reservation not found")
	ErrReservationConflict = errors.New("reservation conflict")

	// rooms
	ErrRoomNotFound   = errors.New("room not found")
	ErrRoomIDRequired = errors.New("room_id is required")
)
