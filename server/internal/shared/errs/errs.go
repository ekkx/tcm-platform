package errs

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code     string
	Message  string
	GRPCCode codes.Code
	Cause    error
}

func New(code, message string, grpcCode codes.Code) *Error {
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

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == t.Code
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

func (e *Error) WithMessage(message string) *Error {
	return &Error{
		Code:     e.Code,
		Message:  message,
		GRPCCode: e.GRPCCode,
		Cause:    e.Cause,
	}
}

// GRPCStatus implements grpc/status.Status interface
func (e *Error) GRPCStatus() *status.Status {
	return status.New(e.GRPCCode, e.Error())
}
