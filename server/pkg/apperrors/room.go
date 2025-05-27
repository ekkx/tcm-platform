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
)
