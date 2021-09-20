package booking

import (
	"context"
	"time"
)

// Unavailability represents a period of time that a resource is unavailable to
// be booked for.
type Unavailability struct {
	ID int `json:"id"`

	// The resource that the unavailability is for.
	ResourceID int       `json:"resourceId"`
	Resource   *Resource `json:"resource"`

	// The time that the unavavailbility begins at.
	StartTime time.Time `json:"startTime"`

	// The time that the unavavailability finishes at.
	EndTime time.Time `json:"endTime"`
}

// UnavailabilityService represents a service for managing unavailabilities.
type UnavailabilityService interface {
	// Retrieves a single unavailability by ID along with the associated resource.
	// Returns ENOTFOUND if unavailability does not exist or user does not have
	// permission to view it.
	FindUnavailabilityByID(ctx context.Context, req FindUnavailabilityByIDRequest) FindUnavailabilityByIDResponse

	// Retreives a list of unavailabilities based on a filter. Only returns
	// unavailabilities that are accessible to the user. Also returns a count of
	// total matching unavailabilities which may be different from the number of
	// returned unavailabilities if the "Limit" field is set.
	FindUnavailabilities(ctx context.Context, req FindUnavailabilitiesRequest) FindUnavailabilitiesResponse

	// Create a new unavailability and assigns the current user as the owner.
	CreateUnavailability(ctx context.Context, req CreateUnavailabilityRequest) CreateUnavailabilityResponse

	// Updates an existing unavailbility by ID. Only the unavailability owner can
	// update an unavailability. Returns the new unavailability state even if
	// there was an error during update.
	//
	// Returns ENOTFOUND if the unavailability does not exist or the user does not
	// have permission to update it.
	UpdateUnavailability(ctx context.Context, req UpdateUnavailabilityRequest) UpdateUnavailabilityResponse

	// Permanently removes a unavailability by ID. Only the unavailability owner
	// may delete a unavailability. Returns ENOTFOUND if the unavailability does
	// not exist or the user does not have permission to delete it.
	DeleteUnavailability(ctx context.Context, req DeleteUnavailabilityRequest) DeleteUnavailabilityResponse
}

// UnavailabilityServiceMiddleware defines a middleware for an unavailability service.
type UnavailabilityServiceMiddleware func(UnavailabilityService) UnavailabilityService

// FindUnavailabilityByIDRequest represents a request used by UnavailabilityService.FindUnavailabilityByID.
type FindUnavailabilityByIDRequest struct {
	ID         int `json:"id" source:"url"`
	ResourceID int `json:"resourceId" source:"resourceId"`
}

// Validate a FindUnavailabilitiesRequest. Returns a ValidationError for each
// requirement that fails.
func (r FindUnavailabilityByIDRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{
			{Name: "id", Reason: "Must be at least 1"},
		}
	}
	return nil
}

// FindUnavailabilityByIDResponse represents a response returned by the
// FindUnavailabilityID method of an UnavailabilityService.
type FindUnavailabilityByIDResponse struct {
	*Unavailability
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindUnavailabilityByIDResponse) Error() error { return r.Err }

// FindUnavailabilitiesRequest represents a query used by
// UnavailabilitiesService.FindUnavailabilities.
type FindUnavailabilitiesRequest struct {
	// Filtering fields.
	ID         *int       `json:"id" source:"query"`
	ResourceID int        `json:"resourceId" source:"query"`
	From       *time.Time `json:"from" source:"query"`
	To         *time.Time `json:"to" source:"query"`

	// Restrict to subset of range.
	Offset int `json:"offset" source:"query"`
	Limit  int `json:"limit" source:"query"`

	// Resource property to order by.
	OrderBy *string `json:"orderBy" source:"query"`
}

// Validate a FindUnavailabilitiesRequest. Returns a ValidationError for each
// requirement that fails.
func (r FindUnavailabilitiesRequest) Validate() []ValidationError {
	errs := make([]ValidationError, 0)
	if r.ID != nil && *r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.ResourceID < 1 {
		errs = append(errs, ValidationError{Name: "resourceId", Reason: "Must be at least 1"})
	}
	if r.From != nil && r.To != nil && r.From.After(*r.To) {
		errs = append(errs, ValidationError{Name: "to", Reason: "Must not be earlier than 'from'"})
	}
	if r.Offset < 0 {
		errs = append(errs, ValidationError{Name: "offset", Reason: "Must be greater than or equal to 0"})
	}
	if r.Limit < 0 {
		errs = append(errs, ValidationError{Name: "limit", Reason: "Must be greater than or equal to 0"})
	}
	validOrderByValues := []string{"id", "resourceId", "startTime", "endTime"}
	if r.OrderBy != nil && !Strings(validOrderByValues).contains(*r.OrderBy) {
		errs = append(errs, ValidationError{Name: "orderBy", Reason: "Must be a valid property name"})
	}
	return errs
}

// FindUnavailabilitiesResponse represents a response returned by the FindUnavailabilities
// method of an UnavailabilitiesService.
type FindUnavailabilitiesResponse struct {
	Unavailabilities []*Unavailability `json:"unavailabilities,omitempty"`
	TotalItems       int               `json:"totalItems"`
	Err              error             `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindUnavailabilitiesResponse) Error() error { return r.Err }

// CreateUnavailability represents a request used by UnavailabilityService.CreateUnavailability.
type CreateUnavailabilityRequest struct {
	ResourceID int       `json:"resourceId" source:"json"`
	StartTime  time.Time `json:"startTime" source:"json"`
	EndTime    time.Time `json:"endTime" source:"json"`
}

// Validate a CreateUnavailabilityRequest. Returns a ValidationError for each
// requirement that fails.
func (r CreateUnavailabilityRequest) Validate() []ValidationError {
	var errs []ValidationError
	if r.ResourceID < 1 {
		errs = append(errs, ValidationError{Name: "resourceId", Reason: "Must be at least 1"})
	}
	if r.StartTime.After(r.EndTime) {
		errs = append(errs, ValidationError{Name: "endTime", Reason: "Must not be earlier than 'startTime'"})
	}
	return errs
}

// CreateUnavailabilityResponse represents a response returned by the CreateUnavailability
// method of a UnavailabilityService.
type CreateUnavailabilityResponse struct {
	*Unavailability
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateUnavailabilityResponse) Error() error { return r.Err }

// UpdateUnavailabilityRequest is a request used by UnavailabilityService.UpdateUnavailability.
type UpdateUnavailabilityRequest struct {
	ID         int       `json:"id" source:"url"`
	ResourceID int       `json:"resourceId" source:"json"`
	StartTime  time.Time `json:"startTime" source:"startTime"`
	EndTime    time.Time `json:"endTime" source:"endTime"`
}

// Validate an UpdateUnavailabilityRequest. Returns a ValidationError for each
// requirement that fails.
func (r UpdateUnavailabilityRequest) Validate() []ValidationError {
	var errs []ValidationError
	if r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.ResourceID < 1 {
		errs = append(errs, ValidationError{Name: "resourceId", Reason: "Must be at least 1"})
	}
	if r.StartTime.After(r.EndTime) {
		errs = append(errs, ValidationError{Name: "endTime", Reason: "Must not be earlier than 'startTime'"})
	}
	return errs
}

// UpdateUnavailabilityResponse represents a response returned by the UpdateUnavailability
// method of an UnavailabilityService.
type UpdateUnavailabilityResponse struct {
	*Unavailability
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateUnavailabilityResponse) Error() error { return r.Err }

// DeleteUnavailabilityRequest represents a request used by UnavailabilityService.DeleteUnavailabilityRequest.
type DeleteUnavailabilityRequest struct {
	ID         int `json:"id" source:"id"`
	ResourceID int `json:"resourceId" source:"resourceId"`
}

// Validate a DeleteUnavailabilityRequest. Returns a ValidationError for each
// requirement that fails.
func (r DeleteUnavailabilityRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{{Name: "id", Reason: "Must be at least 1"}}
	}
	return nil
}

// DeleteUnavailabilityResponse represents a response returned by the DeleteUnavailability
// method of a UnavailabilityService.
type DeleteUnavailabilityResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteUnavailabilityResponse) Error() error { return r.Err }

// UnavailabilityValidationMiddleware represents a middleware. for validating requests passed to a
// unavailability service.
func UnavailabilityValidationMiddleware() UnavailabilityServiceMiddleware {
	return func(next UnavailabilityService) UnavailabilityService {
		return unavailabilityValidationMiddleware{next}
	}
}

type unavailabilityValidationMiddleware struct {
	UnavailabilityService
}

// FindUnavailabilityByID validates a FindUnavailabilityByID request made to a UnavailabilityService.
func (mw unavailabilityValidationMiddleware) FindUnavailabilityByID(ctx context.Context, req FindUnavailabilityByIDRequest) FindUnavailabilityByIDResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindUnavailabilityByIDResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.FindUnavailabilityByID(ctx, req)
}

// FindUnavailabilities validates a FindUnavailabilitysRequest. Forwards the request to the
// next middleware or service if valid and returns an error otherwise.
func (mw unavailabilityValidationMiddleware) FindUnavailabilities(ctx context.Context, req FindUnavailabilitiesRequest) FindUnavailabilitiesResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindUnavailabilitiesResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.FindUnavailabilities(ctx, req)
}

// CreateUnavailability validates a CreateUnavailabilityRequest. Forwards the request to the
// next middleware or service if the request is valid and returns an error otherwise.
func (mw unavailabilityValidationMiddleware) CreateUnavailability(ctx context.Context, req CreateUnavailabilityRequest) CreateUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateUnavailabilityResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.CreateUnavailability(ctx, req)
}

// UpdateUnavailability validates an UpdateUnavailabilityRequest. Forwards the request to
// the next middleware or service if the request is valid and returns an error
// otherwise.
func (mw unavailabilityValidationMiddleware) UpdateUnavailability(ctx context.Context, req UpdateUnavailabilityRequest) UpdateUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return UpdateUnavailabilityResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.UpdateUnavailability(ctx, req)
}

// DeleteUnavailability validates a DeleteUnavailabilityRequest. Forwards the request to the
// next middleware or service if the request is valid and returns an error
// otherwise.
func (mw unavailabilityValidationMiddleware) DeleteUnavailability(ctx context.Context, req DeleteUnavailabilityRequest) DeleteUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return DeleteUnavailabilityResponse{Err: WrapValidationErrors(errs)}
	}

	return mw.UnavailabilityService.DeleteUnavailability(ctx, req)
}
