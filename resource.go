package booking

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// Resource represents a bookable item.
type Resource struct {
	ID int `json:"id"`

	// Organization that the resource belongs to.
	OrganizationID int           `json:"organizationId"`
	Organization   *Organization `json:"organization,omitempty"`

	// The name given to the resource.
	Name string `json:"name"`

	// A description of what the resource is.
	Description string `json:"description"`

	// A collection of slots that are valid for the resource. These
	// indicate when a resource can be booked.
	Slots []*Slot `json:"slots"`

	// The timezone of the resource as a UTC offset. e.g. UTC+00:00
	Timezone string `json:"timezone"`

	// A password to protect the resource. Used to prevent bookings from being
	// made by unauthorized users.
	Password string `json:"password"`

	// The price of the resource to the customer
	Price int `json:"price"`

	// The upfront price that needs to be paid by the customer in order to make a booking.
	BookingPrice int `json:"bookingPrice"`

	// Timestamps for booking creation and last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Slot represents an available period of time of a resource
type Slot struct {
	// The day that the slot is for.
	Day string `json:"day"`

	// The time that the slot begins at.
	StartTime string `json:"startTime"`

	// The time that the slot ends at.
	EndTime string `json:"endTime"`

	// The number of bookings that can be made against the slot on the same day.
	// Nil if there is no limit to the number of bookings that can be made against
	// the slot.
	Quantity *int `json:"quantity"`
}

// ResourceService represents a service for managing resources.
type ResourceService interface {
	// FindResourceByID retrieves a single resource by ID along with associated availabilities.
	// Returns ENOTFOUND if resource does not exist or user does not have
	// permission to view it.
	FindResourceByID(ctx context.Context, req FindResourceByIDRequest) (*Resource, error)

	// FindResources retrieves a lit of resources based on a filter. Only returns
	// resources that accessible to the user. Also returns a count of total matching bookings
	// which may be different from the number of returned bookings if the "Limit" field is set.
	FindResources(ctx context.Context, req FindResourcesRequest) ([]*Resource, int, error)

	// CreateResource creates a new resource and assigns the current user as the owner.
	// Returns the created Resource.
	CreateResource(ctx context.Context, req CreateResourceRequest) (*Resource, error)

	// UpdateResource updates an existing resource by ID. Only the resource owner can update a
	// resource. Returns the new resource state even if there was an error during update.
	//
	// Returns ENOTFOUND if the resource does not exist or the user does not have
	// permission to update it.
	UpdateResource(ctx context.Context, req UpdateResourceRequest) (*Resource, error)

	// DeleteResource permanently removes a resource by ID. Only the resource owner may delete a
	// resource. Returns ENOTFOUND if the resource does not exist or the user does not have
	// permission to delete it.
	DeleteResource(ctx context.Context, req DeleteResourceRequest) error
}

// FindResourceByIDRequest represents a request used by ResourceService.FindResourceByID.
type FindResourceByIDRequest struct {
	ID int `json:"id"`
	// Expand []string `json:"expand"`
}

// Validate a FindResourceByIDRequest. Returns validation.Errors if validation
// fails.
func (r FindResourceByIDRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{
			{Name: "id", Reason: "Must be at least 1"},
		}
	}
	return nil
}

// FindResourceByIDResponse represents a response returned by the
// FindResourceByID method of a ResourceService.
type FindResourceByIDResponse struct {
	*Resource
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindResourceByIDResponse) Error() error { return r.Err }

// FindResourcesRequest represents a query used by
// ResourceRequestHandler.HandleFindResourcesQuery.
type FindResourcesRequest struct {
	// Filtering fields.
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Booking property to order by.
	OrderBy *string `json:"orderBy"`
}

// Validate a FindResourcesRequest. Returns validation.Errors if validation
// fails.
func (r FindResourcesRequest) Validate() []ValidationError {
	errs := make([]ValidationError, 0)
	if r.ID != nil && *r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.Offset < 0 {
		errs = append(errs, ValidationError{Name: "offset", Reason: "Must be greater than or equal to 0"})
	}
	if r.Limit < 0 {
		errs = append(errs, ValidationError{Name: "limit", Reason: "Must be greater than or equal to 0"})
	}
	validOrderByValues := []string{"name", "description", "timezone", "password", "price", "bookingPrice", "createdAt", "updatedAt"}
	if r.OrderBy != nil && !Strings(validOrderByValues).contains(*r.OrderBy) {
		errs = append(errs, ValidationError{Name: "orderBy", Reason: "Must be a valid property name"})
	}
	return errs
}

// FindResourcesResponse represents a response returned by the FindResources
// method of a ResourceService.
type FindResourcesResponse struct {
	Resources  []*Resource `json:"resources,omitempty"`
	TotalItems int         `json:"totalItems"`
	Err        error       `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindResourcesResponse) Error() error { return r.Err }

// CreateResourceRequest represents a request used by ResourceService.CreateResource.
type CreateResourceRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Slots        []*Slot `json:"slots"`
	Timezone     string  `json:"timezone"`
	Password     string  `json:"password"`
	Price        int     `json:"price"`
	BookingPrice int     `json:"bookingPrice"`
}

// Validate a CreateResourceRequest. Returns validation.Errors if validation
// fails
func (r CreateResourceRequest) Validate() []ValidationError {
	var errs []ValidationError
	if r.Name == "" {
		errs = append(errs, ValidationError{Name: "name", Reason: "Name is required"})
	}
	if len(r.Slots) == 0 {
		errs = append(errs, ValidationError{Name: "slots", Reason: "Must contain at least one slot"})
	}
	if r.Price < 0 {
		errs = append(errs, ValidationError{Name: "price", Reason: "Cannot be less than 0"})
	}
	if r.BookingPrice < 0 {
		errs = append(errs, ValidationError{Name: "bookingPrice", Reason: "Cannot be less than 0"})
	}
	if !validTimezone(r.Timezone) {
		errs = append(errs, ValidationError{Name: "timezone", Reason: "Must be valid timezone in the format UTC±HH:MM"})
	}
	errs = append(errs, validateSlots(r.Slots)...)
	return errs
}

// CreateResourceResponse represents a response returned by the CreateResource
// method of a ResourceService.
type CreateResourceResponse struct {
	*Resource
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateResourceResponse) Error() error { return r.Err }

// UpdateResourceByIDRequest represents a request used by ResourceService.UpdateResourceRequest.
type UpdateResourceRequest struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Timezone     string  `json:"timezone"`
	Password     string  `json:"password"`
	Price        int     `json:"price"`
	BookingPrice int     `json:"bookingPrice"`
	Slots        []*Slot `json:"slots"`
}

// Validate an UpdateResourceRequest. Returns validation.Errors if validation
// fails.
func (r UpdateResourceRequest) Validate() []ValidationError {
	var errs []ValidationError
	if r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.Name == "" {
		errs = append(errs, ValidationError{Name: "name", Reason: "Name is required"})
	}
	if len(r.Slots) == 0 {
		errs = append(errs, ValidationError{Name: "slots", Reason: "Must contain at least one slot"})
	}
	if r.Price < 0 {
		errs = append(errs, ValidationError{Name: "price", Reason: "Cannot be less than 0"})
	}
	if r.BookingPrice < 0 {
		errs = append(errs, ValidationError{Name: "bookingPrice", Reason: "Cannot be less than 0"})
	}
	if !validTimezone(r.Timezone) {
		errs = append(errs, ValidationError{Name: "timezone", Reason: "Must be valid timezone in the format UTC±HH:MM"})
	}
	errs = append(errs, validateSlots(r.Slots)...)
	return errs
}

// UpdateResourcesResponse represents a response returned by the UpdateResource
// method of a ResourceService.
type UpdateResourceResponse struct {
	*Resource
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateResourceResponse) Error() error { return r.Err }

// DeleteResourceRequest represents a request used by ResourceService.DeleteResourceRequest.
type DeleteResourceRequest struct {
	ID int `json:"id"`
}

// Validate a DeleteResourceRequest. Returns validation.Errors if validation
// fails.
func (r DeleteResourceRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{{Name: "id", Reason: "Must be at least 1"}}
	}
	return nil
}

// DeleteResourceResponse represents a response returned by the DeleteResource
// method of a ResourceService.
type DeleteResourceResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteResourceResponse) Error() error { return r.Err }

// ResourceServiceMiddleware defines a middleware for a resource request handler.
type ResourceServiceMiddleware func(service ResourceService) ResourceService

// ResourceValidationMiddleware represents a middleware. for validating requests passed to a
// resource service.
func ResourceValidationMiddleware() ResourceServiceMiddleware {
	return func(next ResourceService) ResourceService {
		return resourceValidationMiddleware{next}
	}
}

type resourceValidationMiddleware struct {
	ResourceService
}

// processValidation checks if the error is of type validation.Errors and if it
// is wraps them in a domain Error. If not then it is expected that the error
// should already be a domain Error and the value is returned as is.
func wrapValidationErrors(errs []ValidationError) error {
	return Error{
		Code:   EINVALID,
		Detail: "One or more validation errors occurred while processing your request.",
		Title:  "Invalid request",
		Params: errs,
	}
}

// FindResourceByID validates a FindResourceByID request made to a ResourceService.
func (mw resourceValidationMiddleware) FindResourceByID(ctx context.Context, req FindResourceByIDRequest) (*Resource, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return nil, wrapValidationErrors(errs)
	}
	return mw.ResourceService.FindResourceByID(ctx, req)
}

// FindResources validates a FindResourcesRequest. Forwards the request to the
// next middleware or service if valid and returns an error otherwise.
func (mw resourceValidationMiddleware) FindResources(ctx context.Context, req FindResourcesRequest) ([]*Resource, int, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return nil, 0, wrapValidationErrors(errs)
	}
	return mw.ResourceService.FindResources(ctx, req)
}

// CreateResource validates a CreateResourceRequest. Forwards the request to the
// next middleware or service if the request is valid and returns an error otherwise.
func (mw resourceValidationMiddleware) CreateResource(ctx context.Context, req CreateResourceRequest) (*Resource, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return nil, wrapValidationErrors(errs)
	}
	return mw.ResourceService.CreateResource(ctx, req)
}

// UpdateResource validates an UpdateResourceRequest. Forwards the request to
// the next middleware or service if the request is valid and returns an error
// otherwise.
func (mw resourceValidationMiddleware) UpdateResource(ctx context.Context, req UpdateResourceRequest) (*Resource, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return nil, wrapValidationErrors(errs)
	}
	return mw.ResourceService.UpdateResource(ctx, req)
}

// DeleteResource validates a DeleteResourceRequest. Forwards the request to the
// next middleware or service if the request is valid and returns an error
// otherwise.
func (mw resourceValidationMiddleware) DeleteResource(ctx context.Context, req DeleteResourceRequest) error {
	errs := req.Validate()
	if len(errs) > 0 {
		return wrapValidationErrors(errs)
	}

	return mw.ResourceService.DeleteResource(ctx, req)
}

// validateSlots checks the validity of a []*Slot. Returns a []ValidationError
// containing a ValidationError for each issue found. Checks that startTime and
// endTimes are in the correct format, that startTime is earlier than endTime,
// and that there are no overlapping slots found for any given day.
func validateSlots(slots []*Slot) []ValidationError {
	var errParams []ValidationError
	// Validate slots
	// Check that they all have correct time format
	for i, s := range slots {
		timesAreValid := true
		err := validateSlotTime(s.StartTime)
		if err != nil {
			errParams = append(errParams, ValidationError{
				Name:   fmt.Sprintf("slots[%d].startTime", i),
				Reason: "Must be a valid time in the format HH:MM",
			})
			timesAreValid = false
		}
		err = validateSlotTime(s.EndTime)
		if err != nil {
			errParams = append(errParams, ValidationError{
				Name:   fmt.Sprintf("slots[%d].endTime", i),
				Reason: "Must be a valid time in the format HH:MM",
			})
			timesAreValid = false
		}
		if !timesAreValid {
			continue
		}
		startPrecedesEnd, err := timePrecedes(s.StartTime, s.EndTime)
		if err != nil || !startPrecedesEnd {
			errParams = append(errParams, ValidationError{
				Name:   fmt.Sprintf("slots[%d]", i),
				Reason: "Start time must be earlier than end time",
			})
		}
	}

	// Group slots by day
	daySlots := make(map[string][]*Slot)

	for i, s := range slots {
		daySlots[s.Day] = append(daySlots[s.Day], slots[i])
	}

	// Ensure that slot start and end times are valid for all days.
	for k, v := range daySlots {
		// Ensure slots do not overlap
		// Sort slots by start time
		sort.SliceStable(v, func(i, j int) bool {
			timeI, err := time.Parse("15:04", v[i].StartTime)
			if err != nil {
				return false
			}
			timeJ, err := time.Parse("15:04", v[j].StartTime)
			if err != nil {
				return false
			}
			return timeI.Before(timeJ)
		})

		for i := range v {
			if i == len(v)-1 {
				continue
			}
			precedes, err := timePrecedes(v[i].EndTime, v[i+1].StartTime)
			if err != nil {
				continue
			}
			// If next slot start time is earlier than current slot end time then add a
			// validation error.
			if !precedes {
				equal, err := timeEqual(v[i].EndTime, v[i+1].StartTime)
				if err != nil {
					continue
				}
				if !equal {
					errParams = append(errParams, ValidationError{
						Name:   "slots",
						Reason: fmt.Sprintf("Overlapping start and end times detected for slots with day '%s'", k),
					})
				}
			}
		}
	}
	return errParams
}
