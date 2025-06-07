package errs

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
	ErrInvalidTimeRange = &Error{
		Code:     "invalid_time_range",
		Message:  "invalid time range",
		GRPCCode: codes.InvalidArgument,
	}
	ErrDateMustBeTodayOrFuture = &Error{
		Code:     "date_must_be_today_or_future",
		Message:  "date must be today or future",
		GRPCCode: codes.InvalidArgument,
	}
)
