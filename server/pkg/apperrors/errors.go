package apperrors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

type Error struct {
	Code     string
	Message  string
	GRPCCode codes.Code
	Cause    error
}

func New(code, message string, grpcCode codes.Code, httpStatus int) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		GRPCCode: grpcCode,
	}
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func Is(err, target error) bool {
	if err == nil || target == nil {
		return false
	}

	for {
		if err == target {
			return true
		}

		type isWrapper interface {
			Is(error) bool
		}
		if x, ok := err.(isWrapper); ok && x.Is(target) {
			return true
		}

		err = errors.Unwrap(err)
		if err == nil {
			return false
		}
	}
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) WithCause(cause error) *Error {
	return &Error{
		Code:     e.Code,
		Message:  e.Message,
		GRPCCode: e.GRPCCode,
		Cause:    cause,
	}
}

var (
	// auth errors
	ErrUnauthorized = &Error{
		Code:     "unauthorized",
		Message:  "unauthorized",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidToken = &Error{
		Code:     "invalid_token",
		Message:  "invalid token",
		GRPCCode: codes.Unauthenticated,
	}
	ErrTokenExpired = &Error{
		Code:     "token_expired",
		Message:  "token expired",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidTokenScope = &Error{
		Code:     "invalid_token_scope",
		Message:  "invalid token scope",
		GRPCCode: codes.PermissionDenied,
	}
	// user errors
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
	// reservation errors
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
	// room errors
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
	// internal errors
	ErrInternal = &Error{
		Code:     "internal_error",
		Message:  "internal server error",
		GRPCCode: codes.Internal,
	}
)
