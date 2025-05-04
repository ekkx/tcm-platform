package apperrors

import (
	"errors"
	"net/http"
)

var errorCodeMap = map[error]int{
	// -- auth --

	ErrUnauthorized:      -110000,
	ErrInvalidToken:      -110001,
	ErrTokenExpired:      -110002,
	ErrInvalidTokenScope: -110003,

	// -- users --

	ErrUserNotFound:        -120000,
	ErrRequestUserNotFound: -120001,

	// -- reservations --

	ErrReservationNotFound: -130000,
	ErrReservationConflict: -130001,

	// -- rooms --

	ErrRoomNotFound:   -140000,
	ErrRoomIDRequired: -140001,
}

func getErrorCode(err error) int {
	for e, code := range errorCodeMap {
		if errors.Is(err, e) {
			return code
		}
	}
	return http.StatusInternalServerError
}
