package errs

import "connectrpc.com/connect"

var (
	ErrRoomNotFound = &Error{
		Code:        "room_not_found",
		Message:     "room not found",
		ConnectCode: connect.CodeNotFound,
	}
	ErrRoomIDRequired = &Error{
		Code:        "room_id_required",
		Message:     "room_id is required",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrReservationDateRequired = &Error{
		Code:        "reservation_date_required",
		Message:     "reservation date is required",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrNoAvailableRoom = &Error{
		Code:        "no_available_room",
		Message:     "no available room for the specified time",
		ConnectCode: connect.CodeNotFound,
	}
)
