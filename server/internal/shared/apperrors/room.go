package apperrors

import "google.golang.org/grpc/codes"

var (
	ErrRoomNotFound = &Error{
		Code:     "room_not_found",
		Message:  "room not found",
		GRPCCode: codes.NotFound,
	}
	ErrRoomIDRequired = &Error{
		Code:     "room_id_required",
		Message:  "room_id is required",
		GRPCCode: codes.InvalidArgument,
	}
	ErrReservationDateRequired = &Error{
		Code:     "reservation_date_required",
		Message:  "reservation date is required",
		GRPCCode: codes.InvalidArgument,
	}
	ErrNoAvailableRoom = &Error{
		Code:     "no_available_room",
		Message:  "no available room for the specified time",
		GRPCCode: codes.NotFound,
	}
)
