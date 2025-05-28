package apperrors

import "google.golang.org/grpc/codes"

var (
	ErrReservationNotFound = &Error{
		Code:     "reservation_not_found",
		Message:  "reservation not found",
		GRPCCode: codes.NotFound,
	}
	ErrReservationConflict = &Error{
		Code:     "reservation_conflict",
		Message:  "reservation conflict",
		GRPCCode: codes.AlreadyExists,
	}
	ErrNotYourReservation = &Error{
		Code:     "not_your_reservation",
		Message:  "not your reservation",
		GRPCCode: codes.PermissionDenied,
	}
)
