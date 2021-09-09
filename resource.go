package booking

import (
	"context"
	"time"
)

// Resource represents a bookable item.
type Resource struct {
	ID int `json:"id"`

	// Owner of the resource. Only the owner of the resource may make changes to
	// it.
	UserID int   `json:"userId"`
	User   *User `json:"user"`

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
	// Retrieves a single resource by ID along with associated availabilities.
	// Returns ENOTFOUND if resource does not exist or user does not have
	// permission to view it.
	FindResourceByID(ctx context.Context, id int) (*Resource, error)

	// Retreives a lit of resources based on a filter. Only returns resources that
	// accessible to the user. Also returns a count of total matching bookings
	// which may be different from the number of returned bookings if the "Limit"
	// field is set.
	FindResources(ctx context.Context, filter ResourceFilter) ([]*Resource, int, error)

	// Creates a new resource and assigns the current user as the owner.
	CreateResource(ctx context.Context, resource *Resource) error

	// Updates an existing resource by ID. Only the resource owner can update a
	// resource. Returns the new resource state even if there was an error during
	// update.
	//
	// Returns ENOTFOUND if the resource does not exist or the user does not have
	// permission to update it.
	UpdateResource(ctx context.Context, id int, upd ResourceUpdate) (*Booking, error)

	// Permanently removes a resource by ID. Only the resource owner may delete a
	// resource. Returns ENOTFOUND if the resource does not exist or the user does
	// not have permission to delete it.
	DeleteResource(ctx context.Context, id int) error
}

// ResourceFilter represents a filter used by FindResources()
type ResourceFilter struct {
	// Filtering fields.
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Booking property to order by.
	OrderBy string `json:"orderBy"`
}

// ResourceUpdate represents a set of fields to update on a resource via UpdateResource().
type ResourceUpdate struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Timezone     string  `json:"timezone"`
	Password     string  `json:"password"`
	Price        int     `json:"price"`
	BookingPrice int     `json:"bookingPrice"`
	Slots        []*Slot `json:"slots"`
}
