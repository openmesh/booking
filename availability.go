package booking

import (
	"context"
	"time"
)

// Availability represents a slot that a resource is available with a specific
// date and time associated with it.
type Availability struct {
	// The resource that the availability is for.
	ResourceID int       `json:"resourceId"`
	Resource   *Resource `json:"resource"`

	// Information about the time of the availability.
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// AvailabilityService represents a service for querying resource availability.
type AvailabilityService interface {
	FindAvailabilties(ctx context.Context, filter AvailabilityFilter) ([]*Availability, int, error)
}

// AvailabilityFilter represents a filter used by FindAvailabilities()
type AvailabilityFilter struct {
	// Filtering fields.
	ResourceID *int `json:"resourceId"`

	StartTimeAfter time.Time `json:"startTimeAfter"`
	EndTimeBefore  time.Time `json:"endTimeBefore"`

	// Restrict to subset of range.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// Availability property to order by.
	OrderBy string `json:"orderBy"`
}

// AvailabilityServiceMiddleware defines a middleware for an availability service.
type AvailabilityServiceMiddleware func(AvailabilityService) AvailabilityService
