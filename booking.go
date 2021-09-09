package booking

import (
	"context"
	"time"
)

// Booking represents a booking of a resource that has been made by a customer.
type Booking struct {
	ID int `json:"id"`

	// Owner of the booking. Only the owner of the booking may make changes to it.
	UserID int   `json:"userId"`
	User   *User `json:"user"`

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
	FindBookingByID(ctx context.Context, id int) (*Booking, error)

	// Retreives a lit of bookings based on a filter. Only returns bookings that
	// are accessible to the user. Also returns a count of total matching bookings
	// which may be different from the number of returned bookings if the "Limit"
	// field is set.
	FindBookings(ctx context.Context, filter BookingFilter) ([]*Booking, int, error)

	// Creates a new booking and assigns the current user as the owner.
	CreateBooking(ctx context.Context, booking *Booking) error

	// Updates an existing booking by ID. Only the booking owner can update a
	// booking. Returns the new booking state even if there was an error during
	// update.
	//
	// Returns ENOTFOUND if the booking does not exist or the user does not have
	// permission to update it.
	UpdateBooking(ctx context.Context, id int, upd BookingUpdate) (*Booking, error)

	// Permanently removes a booking by ID. Only the booking owner may delete a
	// booking. Returns ENOTFOUND if the booking does not exist or the user does
	// not have permission to delete it.
	DeleteBooking(ctx context.Context, id int) error
}

// BookingFilter represents a filter used by FindBookings()
type BookingFilter struct {
	// Filtering fields.
	ID             *int      `json:"id"`
	ResourceID     *int      `json:"resourceId"`
	Status         *string   `json:"status"`
	StartTimeAfter time.Time `json:"startTimeAfter"`
	EndTimeAfter   time.Time `json:"endTimeAfter"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Booking property to order by.
	OrderBy string `json:"orderBy"`
}

// BookingUpdate represents a set of fields to update on a booking.
type BookingUpdate struct {
	Metadata map[string]string `json:"metadata"`
	Status   string            `json:"status"`
}
