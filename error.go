package booking

import (
	"errors"
	"fmt"
)

// Application error codes.
//
// NOTE: These are meant to be generic and they map well to HTTP error codes.
// Different applications can have very different error code requirements so
// these should be expanded as needed (or introduce subcodes).
const (
	ECONFLICT       = "conflict"
	EINTERNAL       = "internal"
	EINVALID        = "invalid"
	ENOTFOUND       = "not_found"
	ENOTIMPLEMENTED = "not_implemented"
	EUNAUTHORIZED   = "unauthorized"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.
//
// Any non-application error (such as a disk error) should be reported as an
// EINTERNAL error and the human user should only see "Internal error" as the
// message. These low-level internal error details should only be logged and
// reported to the operator of the application (not the end user).
type Error struct {
	// Machine-readable error code.
	Code string `json:"code"`

	// A human-readable description of the specific error.
	Detail string `json:"detail"`

	// A short, human-readable title for the general error type.
	Title string `json:"title"`

	// An optional slice of parameter specific errors. Useful when returning
	// validation errors.
	Params []ErrorParam `json:"params,omitempty"`
}

// ErrorParam is used to return parameter specific issues within an Error.
type ErrorParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

// Error implements the error interface. Not used by the application otherwise.
func (e Error) Error() string {
	return fmt.Sprintf("booking error: code=%s title=%s detail=%s", e.Code, e.Title, e.Detail)
}

// ErrorCode unwraps an application error and returns its code. Non-application
// errors always return EINTERNAL.
func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}
	return EINTERNAL
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "Internal error".
func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Detail
	}
	return "Internal error."
}

// Errorf is a helper function to return an Error with a given code and
// formatted message.
func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:   code,
		Detail: fmt.Sprintf(format, args...),
	}
}

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type Errorer interface {
	Error() error
}
