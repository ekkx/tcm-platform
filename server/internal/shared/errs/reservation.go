package errs

import "connectrpc.com/connect"

var (
	ErrReservationNotFound = &Error{
		Code:        "reservation_not_found",
		Message:     "reservation not found",
		ConnectCode: connect.CodeNotFound,
	}
	ErrReservationConflict = &Error{
		Code:        "reservation_conflict",
		Message:     "reservation conflict",
		ConnectCode: connect.CodeAlreadyExists,
	}
	ErrNotYourReservation = &Error{
		Code:        "not_your_reservation",
		Message:     "not your reservation",
		ConnectCode: connect.CodePermissionDenied,
	}
	ErrInvalidTimeRange = &Error{
		Code:        "invalid_time_range",
		Message:     "invalid time range",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrInvalidDate = &Error{
		Code:        "invalid_date",
		Message:     "invalid date",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrDateMustBeTodayOrFuture = &Error{
		Code:        "date_must_be_today_or_future",
		Message:     "date must be today or future",
		ConnectCode: connect.CodeInvalidArgument,
	}
)
