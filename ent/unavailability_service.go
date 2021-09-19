package ent

import (
	"context"
	"fmt"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/unavailability"
)

type unavailabilityService struct {
	client *Client
}

func NewUnavailabilityService(client *Client) *unavailabilityService {
	return &unavailabilityService{
		client,
	}
}

// Retrieves a single unavailability by ID along with the associated resource.
// Returns ENOTFOUND if unavailability does not exist or user does not have
// permission to view it.
func (s *unavailabilityService) FindUnavailabilityByID(
	ctx context.Context,
	req booking.FindUnavailabilityByIDRequest,
) booking.FindUnavailabilityByIDResponse {
	u, err := s.client.Unavailability.
		Query().
		Where(
			unavailability.ID(req.ID),
			unavailability.ResourceId(req.ResourceID),
			unavailability.HasResourceWith(resource.OrganizationId(booking.OrganizationIDFromContext(ctx))),
		).
		WithResource(func(q *ResourceQuery) {
			q.WithSlots()
		}).
		First(ctx)

	if _, ok := err.(*NotFoundError); ok {
		return booking.FindUnavailabilityByIDResponse{
			Err: booking.Error{
				Code: booking.ENOTFOUND,
				Detail: fmt.Sprintf(
					"'Unavailability' with ID '%d' belonging to a resource with ID '%d' could not be found",
					req.ID,
					req.ResourceID,
				),
				Title: "Unavailability not found",
			},
		}
	}
	if err != nil {
		return booking.FindUnavailabilityByIDResponse{Err: err}
	}

	return booking.FindUnavailabilityByIDResponse{Unavailability: u.toModel()}
}

// Retreives a list of unavailabilities based on a filter. Only returns
// unavailabilities that are accessible to the user. Also returns a count of
// total matching unavailabilities which may be different from the number of
// returned unavailabilities if the "Limit" field is set.
func (s *unavailabilityService) FindUnavailabilities(
	ctx context.Context,
	req booking.FindUnavailabilitiesRequest,
) booking.FindUnavailabilitiesResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindUnavailabilitiesResponse{Err: err}
	}
	query := tx.Unavailability.
		Query().
		Where(
			unavailability.HasResourceWith(resource.OrganizationId(booking.OrganizationIDFromContext(ctx))),
			unavailability.ResourceId(req.ResourceID),
		)

	if req.ID != nil {
		query = query.Where(unavailability.ID(*req.ID))
	}
	if req.From != nil {
		query = query.Where(unavailability.StartTimeGTE(*req.From))
	}
	if req.To != nil {
		query = query.Where(unavailability.EndTimeLTE(*req.To))
	}
	totalItems, err := query.Count(ctx)
	if err != nil {
		return booking.FindUnavailabilitiesResponse{Err: err}
	}
	query = query.Offset(req.Offset)
	if req.Limit == 0 {
		query = query.Limit(10)
	} else {
		query = query.Limit(req.Limit)
	}
	u, err := query.WithResource(func(q *ResourceQuery) {
		q.WithSlots()
	}).All(ctx)
	if err != nil {
		return booking.FindUnavailabilitiesResponse{Err: err}
	}

	return booking.FindUnavailabilitiesResponse{
		Unavailabilities: Unavailabilities(u).toModels(),
		TotalItems:       totalItems,
	}
}

// Create a new unavailability and assigns the current user as the owner.
func (s *unavailabilityService) CreateUnavailability(
	ctx context.Context,
	req booking.CreateUnavailabilityRequest,
) booking.CreateUnavailabilityResponse {
	panic("not implemented") // TODO: Implement
}

// Updates an existing unavailbility by ID. Only the unavailability owner can
// update an unavailability. Returns the new unavailability state even if
// there was an error during update.
//
// Returns ENOTFOUND if the unavailability does not exist or the user does not
// have permission to update it.
func (s *unavailabilityService) UpdateUnavailability(
	ctx context.Context,
	req booking.UpdateUnavailabilityRequest,
) booking.UpdateUnavailabilityResponse {
	panic("not implemented") // TODO: Implement
}

// Permanently removes a unavailability by ID. Only the unavailability owner
// may delete a unavailability. Returns ENOTFOUND if the unavailability does
// not exist or the user does not have permission to delete it.
func (s *unavailabilityService) DeleteUnavailability(
	ctx context.Context,
	req booking.DeleteUnavailabilityRequest,
) booking.DeleteUnavailabilityResponse {
	panic("not implemented") // TODO: Implement
}

func (u *Unavailability) toModel() *booking.Unavailability {
	result := &booking.Unavailability{
		ID:         u.ID,
		ResourceID: u.ResourceId,
		StartTime:  u.StartTime,
		EndTime:    u.EndTime,
	}
	if u.Edges.Resource != nil {
		result.Resource = u.Edges.Resource.toModel()
	}
	return result
}

func (u Unavailabilities) toModels() []*booking.Unavailability {
	var unavailabilities []*booking.Unavailability
	for _, v := range u {
		unavailabilities = append(unavailabilities, v.toModel())
	}
	return unavailabilities
}
