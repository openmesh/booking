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
	ECONFLICT = "conflict"
	// EBOOKINGCONFLICT indicates that a request was made to create a booking that
	// would exceed the quantity available for a resource.
	EBOOKINGCONFLICT = "booking_conflict"
	EINTERNAL        = "internal"
	EINVALID         = "invalid"
	ENOTFOUND        = "not_found"
	// ERESOURCENOTFOUND indicates that a request was made to retrieve a resource
	// that does not exist or is not accessible by the requester.
	ERESOURCENOTFOUND = "resource_not_found"
	// EUNAVAILABILITYNOTFOUND indicates that a request was made to retrieve an
	// unavailability that does not exist or is not accessible by the requester.
	EUNAVAILABILITYNOTFOUND = "unavailability_not_found"
	// EBOOKINGNOTFOUND indicates that a request was made to retrieve a booking
	// that does not exist or is not accessible by the requester.
	EBOOKINGNOTFOUND = "booking_not_found"
	// EUNAVAILABILITYTIMECONFLICT indicates that a request was made to create
	// update an unavailability which would have conflicted with an existing
	// unavailability.
	EUNAVAILABILITYTIMECONFLICT = "unavailability_time_conflict"
	// ERESOURCENAMECONFLICT indicates that a request was to create a resource
	// with a name already taken by an existing resource within the same
	// organiztion.
	ERESOURCENAMECONFLICT = "resource_name_conflict"
	ENOTIMPLEMENTED       = "not_implemented"
	EUNAUTHORIZED         = "unauthorized"
	// EAUTHSOURCENOTCONFIGURED indicates that an attempt was made to use an auth
	// source that had not been set up.
	EAUTHSOURCENOTCONFIGURED = "auth_source_not_configured"
	EAUTHSOURCEUNSUPPORTED   = "auth_source_unsupported"
	EAUTHNOTFOUND            = "auth_not_found"
	EUSERNOTFOUND            = "user_not_found"
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
	Params []ValidationError `json:"params,omitempty"`
}

// Error implements the errorer interface. Not used by the application otherwise.
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

func ValidationErrorf(detail string, errs ...ValidationError) *Error {
	return &Error{
		Code:   EINVALID,
		Detail: "One or more validation errors occurred while processing your request.",
		Title:  "Invalid request",
		Params: errs,
	}
}

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type Errorer interface {
	Error() error
}

// WrapValidationErrors checks if the error is of type validation.Errors and if it
// is wraps them in a domain Error. If not then it is expected that the error
// should already be a domain Error and the value is returned as is.
func WrapValidationErrors(errs []ValidationError) error {
	return Error{
		Code:   EINVALID,
		Detail: "One or more validation errors occurred while processing your request.",
		Title:  "Invalid request",
		Params: errs,
	}
}

// WrapNotFoundError wraps a not
func WrapNotFoundError(entity string) error {
	return Error{
		Code:   ENOTFOUND,
		Detail: fmt.Sprintf("Specified %s could not be found", entity),
		Title:  "Item not found",
	}
}
