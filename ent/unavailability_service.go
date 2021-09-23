package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindUnavailabilityByIDResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	u, err := findUnavailabilityByID(
		ctx,
		tx,
		req.ID,
		func(uq *UnavailabilityQuery) *UnavailabilityQuery {
			uq.WithResource(func(rq *ResourceQuery) {
				rq.WithSlots()
			})
			return uq
		},
	)
	if err != nil {
		return booking.FindUnavailabilityByIDResponse{
			Err: fmt.Errorf("failed to find unavailability: %w", err),
		}
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
		return booking.FindUnavailabilitiesResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	u, totalItems, err := findUnavailabilties(
		ctx,
		tx,
		req,
		func(uq *UnavailabilityQuery) *UnavailabilityQuery {
			uq.WithResource(func(rq *ResourceQuery) {
				rq.WithSlots()
			})
			return uq
		})
	if err != nil {
		return booking.FindUnavailabilitiesResponse{
			Err: fmt.Errorf("failed to find unavailabilities: %w", err),
		}
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
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateUnavailabilityResponse{Err: err}
	}

	err = checkForUnavailabilityTimeConflict(ctx, tx, req.StartTime, req.EndTime)
	if err != nil {
		return booking.CreateUnavailabilityResponse{
			Err: fmt.Errorf("unvailability time conflict check failed: %w", err),
		}
	}

	// Create unavailability
	u, err := tx.Unavailability.
		Create().
		SetResourceID(req.ResourceID).
		SetStartTime(req.StartTime).
		SetEndTime(req.EndTime).
		Save(ctx)
	if err != nil {
		return booking.CreateUnavailabilityResponse{Err: err}
	}

	u.Edges.Resource, err = u.QueryResource().WithSlots().First(ctx)
	if err != nil {
		return booking.CreateUnavailabilityResponse{Err: tx.Rollback()}
	}

	return booking.CreateUnavailabilityResponse{
		Unavailability: u.toModel(),
		Err:            tx.Commit(),
	}
}

func checkForUnavailabilityTimeConflict(ctx context.Context, tx *Tx, st time.Time, et time.Time) error {
	// Ensure that new unavailability will not conflict with any existing availabilities.
	count, err := tx.Unavailability.
		Query().
		Where(
			unavailability.Or(
				// New unavailability begins during an existing unavailability
				unavailability.And(unavailability.StartTimeLTE(st), unavailability.EndTimeGTE(st)),
				// New unavailability ends during an exsting unavailability
				unavailability.And(unavailability.StartTimeLTE(et), unavailability.EndTimeGTE(et)),
				// New unavailability is entirely during an existing unavailability
				unavailability.And(unavailability.StartTimeLTE(st), unavailability.EndTimeGTE(et)),
				// Existing unavailability is entirely during new unavailability
				unavailability.And(unavailability.StartTimeGTE(st), unavailability.EndTimeLTE(et)),
			),
		).
		Count(ctx)
	if err != nil {
		return fmt.Errorf("failed to count unavailabilities")
	}
	if count > 0 {
		return booking.Errorf(
			booking.EUNAVAILABILITYTIMECONFLICT,
			"Specified unavailability would conflict with an already existing unavailability",
		)
	}
	return nil
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
	tx, err := s.client.Tx(ctx)

	orgID := booking.OrganizationIDFromContext(ctx)
	// Retrieve unavailability
	u, err := tx.Unavailability.
		Query().
		Where(
			unavailability.ID(req.ID),
			unavailability.ResourceId(req.ResourceID),
			unavailability.HasResourceWith(resource.OrganizationId(orgID)),
		).First(ctx)

	if _, ok := err.(*NotFoundError); ok {
		return booking.UpdateUnavailabilityResponse{Err: booking.WrapNotFoundError("unavailability")}
	}
	if err != nil {
		return booking.UpdateUnavailabilityResponse{Err: err}
	}

	// Ensure that new unavailability will not conflict with any existing availabilities.
	if count, err := tx.Unavailability.
		Query().
		Where(
			unavailability.IDNEQ(req.ID),
			unavailability.HasResourceWith(resource.ID(req.ResourceID)),
			unavailability.HasResourceWith(resource.OrganizationId(booking.OrganizationIDFromContext(ctx))),
			unavailability.Or(
				// New unavailability begins during an existing unavailability
				unavailability.And(unavailability.StartTimeLTE(req.StartTime), unavailability.EndTimeGTE(req.StartTime)),
				// New unavailability ends during an exsting unavailability
				unavailability.And(unavailability.StartTimeLTE(req.EndTime), unavailability.EndTimeGTE(req.EndTime)),
				// New unavailability is entirely during an existing unavailability
				unavailability.And(unavailability.StartTimeLTE(req.StartTime), unavailability.EndTimeGTE(req.EndTime)),
				// Existing unavailability is entirely during an existing unavailability
				unavailability.And(unavailability.StartTimeGTE(req.StartTime), unavailability.EndTimeLTE(req.EndTime)),
			),
		).
		Count(ctx); err != nil {
		return booking.UpdateUnavailabilityResponse{Err: err}
	} else if count > 0 {
		return booking.UpdateUnavailabilityResponse{
			Err: booking.Error{
				Code:   booking.ECONFLICT,
				Detail: "Request unavailability would conflict with existing unavailabilities. Check start and end times for potential overlap.",
				Title:  "Conflicting unavailability found",
			},
		}
	}

	u, err = tx.Unavailability.UpdateOne(u).SetStartTime(req.StartTime).SetEndTime(req.EndTime).Save(ctx)
	if err != nil {
		return booking.UpdateUnavailabilityResponse{Err: tx.Rollback()}
	}

	u.Edges.Resource, err = u.QueryResource().WithSlots().First(ctx)
	if err != nil {
		return booking.UpdateUnavailabilityResponse{Err: tx.Rollback()}
	}

	// Ensure unavailability exists
	return booking.UpdateUnavailabilityResponse{Unavailability: u.toModel()}
}

// Permanently removes a unavailability by ID. Only the unavailability owner
// may delete a unavailability. Returns ENOTFOUND if the unavailability does
// not exist or the user does not have permission to delete it.
func (s *unavailabilityService) DeleteUnavailability(
	ctx context.Context,
	req booking.DeleteUnavailabilityRequest,
) booking.DeleteUnavailabilityResponse {
	orgID := booking.OrganizationIDFromContext(ctx)

	if affected, err := s.client.Unavailability.
		Delete().
		Where(
			unavailability.ID(req.ID),
			unavailability.ResourceId(req.ResourceID),
			unavailability.HasResourceWith(resource.OrganizationId(orgID))).
		Exec(ctx); err != nil {
		return booking.DeleteUnavailabilityResponse{Err: err}
	} else if affected == 0 {
		return booking.DeleteUnavailabilityResponse{Err: booking.WrapNotFoundError("unavailability")}
	}

	return booking.DeleteUnavailabilityResponse{}
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

// findUnavailabilityByID is a helper function to retrieve an unavailability by
// ID. Returns EUNAVAILABILITYNOTFOUND if the unavailability does not exist or
// is not accessible by the current user.
func findUnavailabilityByID(
	ctx context.Context,
	tx *Tx,
	id int,
	withEdges func(*UnavailabilityQuery) *UnavailabilityQuery,
) (*Unavailability, error) {
	u, err := withEdges(
		tx.Unavailability.
			Query().
			Where(unavailability.ID(id)),
	).First(ctx)

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.EUNAVAILABILITYNOTFOUND, "Could not find unavailability with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

// findUnavailabilties retrieves a slice of resources based on a set of
// parameters. Also returns the total number of records that meet the criteria
// before offsets and limits are applied.
// withEdges exposes the query builder with the intention of allowing the caller
// to eager load edges.
func findUnavailabilties(
	ctx context.Context,
	tx *Tx,
	req booking.FindUnavailabilitiesRequest,
	withEdges func(*UnavailabilityQuery) *UnavailabilityQuery,
) ([]*Unavailability, int, error) {
	query := tx.Unavailability.
		Query().
		Where(unavailability.ResourceId(req.ResourceID))

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
		return nil, 0, fmt.Errorf("failed to count unavailabilities: %w", err)
	}

	query = query.Offset(req.Offset)
	if req.Limit == 0 {
		query = query.Limit(10)
	} else {
		query = query.Limit(req.Limit)
	}
	u, err := withEdges(query).All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query unavailabilities: %w", err)
	}

	return u, totalItems, nil
}
