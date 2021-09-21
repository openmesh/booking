package booking

import (
	"context"
	"time"
)

// Booking represents a booking of a resource that has been made by a customer.
type Booking struct {
	ID int `json:"id"`

	// The resource that the booking has been made for.
	ResourceID int       `json:"resourceId"`
	Resource   *Resource `json:"resource"`

	// Generic information about the booking. Can include things like the
	// customer's personal information.
	Metadata map[string]string `json:"metadata"`

	// The status of the booking.
	Status string `json:"status"`

	// Information about the time of the booking.
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`

	// Timestamps for booking creation and last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BookingService represents a service for managing bookings.
type BookingService interface {
	// Retrieves a single booking by ID along with the associated resource and
	// metadata. Returns ENOTFOUND if booking does not exist or user does not have
	// permission to view it.
	FindBookingByID(ctx context.Context, req FindBookingByIDRequest) FindBookingByIDResponse

	// Retreives a list of bookings based on a filter. Only returns bookings that
	// are accessible to the user. Also returns a count of total matching bookings
	// which may be different from the number of returned bookings if the "Limit"
	// field is set.
	FindBookings(ctx context.Context, req FindBookingsRequest) FindBookingsResponse

	// Creates a new booking and assigns the current user as the owner.
	CreateBooking(ctx context.Context, req CreateBookingRequest) CreateBookingResponse

	// Updates an existing booking by ID. Only the booking owner can update a
	// booking. Returns the new booking state even if there was an error during
	// update.
	//
	// Returns ENOTFOUND if the booking does not exist or the user does not have
	// permission to update it.
	UpdateBooking(ctx context.Context, req UpdateBookingRequest) UpdateBookingResponse

	// Permanently removes a booking by ID. Only the booking owner may delete a
	// booking. Returns ENOTFOUND if the booking does not exist or the user does
	// not have permission to delete it.
	DeleteBooking(ctx context.Context, req DeleteBookingRequest) DeleteBookingResponse
}

// BookingUpdate represents a set of fields to update on a booking.
type BookingUpdate struct {
	ResourceID int               `json:"resourceId"`
	Metadata   map[string]string `json:"metadata"`
	Status     string            `json:"status"`
}

// FindBookingByIDRequest represents a payload used by the FindBookingByID method of a BookingService
type FindBookingByIDRequest struct {
	ID int `json:"id" source:"url"`
}

// Validate a FindBookingByID. Returns a ValidationError for each requirement that fails.
func (r FindBookingByIDRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{
			{Name: "id", Reason: "Must be at least 1"},
		}
	}
	return nil
}

// FindBookingByIDResponse represents a response returned by the FindBookingByID method of a BookingService.
type FindBookingByIDResponse struct {
	*Booking
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindBookingByIDResponse) Error() error { return r.Err }

// FindBookingsRequest represents a payload used by the FindBookings method of a BookingService
type FindBookingsRequest struct {
	// Filtering fields.
	ID             *int       `json:"id" source:"query"`
	ResourceID     *int       `json:"resourceId" source:"query"`
	Status         *string    `json:"status" source:"query"`
	StartTimeAfter *time.Time `json:"startTimeAfter" source:"query"`
	EndTimeBefore  *time.Time `json:"endTimeBefore" source:"query"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Booking property to order by.
	OrderBy string `json:"orderBy"`
}

// Validate a FindBookings. Returns a ValidationError for each requirement that fails.
func (r FindBookingsRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindBookingsResponse represents a response returned by the FindBookings method of a BookingService.
type FindBookingsResponse struct {
	Bookings   []*Booking `json:"bookings,omitempty"`
	TotalItems int        `json:"totalItems"`
	Err        error      `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindBookingsResponse) Error() error { return r.Err }

// CreateBookingRequest represents a payload used by the CreateBooking method of a BookingService
type CreateBookingRequest struct {
	// The resource that the booking has been made for.
	ResourceID int `json:"resourceId" source:"json"`

	// Generic information about the booking. Can include things like the
	// customer's personal information.
	Metadata map[string]string `json:"metadata" source:"json"`

	// The status of the booking.
	Status string `json:"status" source:"json"`

	// Information about the time of the booking.
	StartTime time.Time `json:"startTime" source:"json"`
	EndTime   time.Time `json:"endTime" source:"json"`
}

// Validate a CreateBooking. Returns a ValidationError for each requirement that fails.
func (r CreateBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// CreateBookingResponse represents a response returned by the CreateBooking method of a BookingService.
type CreateBookingResponse struct {
	*Booking
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateBookingResponse) Error() error { return r.Err }

// UpdateBookingRequest represents a payload used by the UpdateBooking method of a BookingService
type UpdateBookingRequest struct {
	ID         int               `json:"id" source:"url"`
	ResourceID int               `json:"resourceId" source:"json"`
	Metadata   map[string]string `json:"metadata" source:"json"`
	Status     string            `json:"status" source:"json"`
	StartTime  time.Time         `json:"startTime" source:"json"`
	EndTime    time.Time         `json:"endTime" source:"json"`
}

// Validate a UpdateBooking. Returns a ValidationError for each requirement that fails.
func (r UpdateBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// UpdateBookingResponse represents a response returned by the UpdateBooking method of a BookingService.
type UpdateBookingResponse struct {
	*Booking
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateBookingResponse) Error() error { return r.Err }

// DeleteBookingRequest represents a payload used by the DeleteBooking method of a BookingService
type DeleteBookingRequest struct {
	ID int `json:"id" source:"url"`
}

// Validate a DeleteBooking. Returns a ValidationError for each requirement that fails.
func (r DeleteBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// DeleteBookingResponse represents a response returned by the DeleteBooking method of a BookingService.
type DeleteBookingResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteBookingResponse) Error() error { return r.Err }

// BookingServiceMiddleware defines a middleware for BookingService
type BookingServiceMiddleware func(service BookingService) BookingService

// BookingValidationMiddleware returns a middleware for validating requests made to a BookingService
func BookingValidationMiddleware() BookingServiceMiddleware {
	return func(next BookingService) BookingService {
		return bookingValidationMiddleware{next}
	}
}

type bookingValidationMiddleware struct {
	BookingService
}

// FindBookingByID validates a FindBookingByIDRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddleware) FindBookingByID(ctx context.Context, req FindBookingByIDRequest) FindBookingByIDResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindBookingByIDResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.BookingService.FindBookingByID(ctx, req)
}

// FindBookings validates a FindBookingsRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddleware) FindBookings(ctx context.Context, req FindBookingsRequest) FindBookingsResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindBookingsResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.BookingService.FindBookings(ctx, req)
}

// CreateBooking validates a CreateBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddleware) CreateBooking(ctx context.Context, req CreateBookingRequest) CreateBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateBookingResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.BookingService.CreateBooking(ctx, req)
}

// UpdateBooking validates a UpdateBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddleware) UpdateBooking(ctx context.Context, req UpdateBookingRequest) UpdateBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return UpdateBookingResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.BookingService.UpdateBooking(ctx, req)
}

// DeleteBooking validates a DeleteBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddleware) DeleteBooking(ctx context.Context, req DeleteBookingRequest) DeleteBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return DeleteBookingResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.BookingService.DeleteBooking(ctx, req)
}
