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
	FindUnavailabilityByID(ctx context.Context, id int) (*Unavailability, error)

	// Retreives a list of unavailabilities based on a filter. Only returns
	// unavailabilities that are accessible to the user. Also returns a count of
	// total matching unavailabilities which may be different from the number of
	// returned unavailabilities if the "Limit" field is set.
	FindUnavailabilities(ctx context.Context, filter UnavailabilityFilter) ([]*Unavailability, int, error)

	// Create a new unavailability and assigns the current user as the owner.
	CreateUnavailability(ctx context.Context, unavailability *Unavailability) error

	// Updates an existing unavailbility by ID. Only the unavailability owner can
	// update an unavailability. Returns the new unavailability state even if
	// there was an error during update.
	//
	// Returns ENOTFOUND if the unavailability does not exist or the user does not
	// have permission to update it.
	UpdateUnavailability(ctx context.Context, id int, upd UnavailabilityUpdate) (*Unavailability, error)

	// Permanently removes a unavailability by ID. Only the unavailability owner
	// may delete a unavailability. Returns ENOTFOUND if the unavailability does
	// not exist or the user does not have permission to delete it.
	DeleteUnavailability(ctx context.Context, id int) error
}

// UnavailabilityFilter represents a filter used by FindUnavailabilities()
type UnavailabilityFilter struct {
	// Filtering fields.
	ID             *int      `json:"id"`
	ResourceID     *int      `json:"resourceId"`
	StartTimeAfter time.Time `json:"startTimeAfter"`
	EndTimeBefore  time.Time `json:"endTimeBefore"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Unavailability property to order by.
	OrderBy string `json:"orderBy"`
}

// UnavailabilityUpdate represents a set of fields to update on a unavailbility.
type UnavailabilityUpdate struct {
	ResourceID int       `json:"resourceId"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

// UnavailabilityServiceMiddleware defines a middleware for an unavailability service.
type UnavailabilityServiceMiddleware func(UnavailabilityService) UnavailabilityService
