package apperrors

import "google.golang.org/grpc/codes"

var (
	ErrUserNotFound = &Error{
		Code:     "user_not_found",
		Message:  "user not found",
		GRPCCode: codes.NotFound,
	}
	ErrRequestUserNotFound = &Error{
		Code:     "request_user_not_found",
		Message:  "request user not found",
		GRPCCode: codes.NotFound,
	}
)
