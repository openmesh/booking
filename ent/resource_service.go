package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/slot"
)

type resourceService struct {
	client *Client
}

func NewResourceService(client *Client) *resourceService {
	return &resourceService{
		client,
	}
}

// FindResourceByID retrieves a single resource by ID along with associated availabilities.
// Returns ERESOURCENOTFOUND if resource does not exist or user does not have
// permission to view it.
func (s *resourceService) FindResourceByID(
	ctx context.Context,
	req booking.FindResourceByIDRequest,
) booking.FindResourceByIDResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindResourceByIDResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}

	r, err := findResourceByID(ctx, tx, req.ID, func(rq *ResourceQuery) *ResourceQuery {
		return rq.WithSlots()
	})
	if err != nil {
		return booking.FindResourceByIDResponse{Err: err}
	}

	r.Edges.Slots, err = r.QuerySlots().All(ctx)
	if err != nil {
		return booking.FindResourceByIDResponse{Err: err}
	}

	return booking.FindResourceByIDResponse{Resource: r.toModel()}
}

// FindResources retrieves a lit of resources based on a filter. Only returns
// resources that accessible to the user. Also returns a count of total matching resources
// which may be different from the number of returned bookings if the "Limit" field is set.
func (s *resourceService) FindResources(
	ctx context.Context,
	req booking.FindResourcesRequest,
) booking.FindResourcesResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindResourcesResponse{Err: fmt.Errorf("failed to start transaction: %w", err)}
	}

	r, totalItems, err := findResources(ctx, tx, req, func(rq *ResourceQuery) *ResourceQuery {
		return rq.WithSlots()
	})
	if err != nil {
		return booking.FindResourcesResponse{Err: fmt.Errorf("failed to find resources: %w", err)}
	}

	return booking.FindResourcesResponse{
		Resources:  Resources(r).toModels(),
		TotalItems: totalItems,
	}
}

// CreateResource creates a new resource and assigns the current user as the owner.
// Returns the created Resource.
func (s *resourceService) CreateResource(
	ctx context.Context,
	req booking.CreateResourceRequest,
) booking.CreateResourceResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateResourceResponse{Err: fmt.Errorf("failed to start transaction: %w", err)}
	}
	defer tx.Rollback()

	err = checkForResourceNameConflict(ctx, tx, req.Name)
	if err != nil {
		return booking.CreateResourceResponse{Err: fmt.Errorf("failed resource name check: %w", err)}
	}

	r, err := createResource(ctx, tx, req)
	if err != nil {
		return booking.CreateResourceResponse{Err: fmt.Errorf("failed to create resource: %w", err)}
	}

	err = tx.Commit()
	if err != nil {
		return booking.CreateResourceResponse{Err: fmt.Errorf("failed to commit transaction: %w", err)}
	}

	return booking.CreateResourceResponse{
		Resource: r.toModel(),
		Err:      nil,
	}
}

// UpdateResource updates an existing resource by ID. Only the resource owner can update a
// resource. Returns the new resource state even if there was an error during update.
//
// Returns ERESOURCENOTFOUND if the resource does not exist or the user does not have
// permission to update it.
func (s *resourceService) UpdateResource(
	ctx context.Context,
	req booking.UpdateResourceRequest,
) booking.UpdateResourceResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.UpdateResourceResponse{Err: fmt.Errorf("failed to start transaction: %w", err)}
	}

	err = checkForResourceNameConflict(ctx, tx, req.Name)
	if err != nil {
		return booking.UpdateResourceResponse{Err: fmt.Errorf("failed resource name check: %w", err)}
	}

	r, err := updateResource(ctx, tx, req)
	if err != nil {
		return booking.UpdateResourceResponse{Err: err}
	}

	err = tx.Commit()
	if err != nil {
		return booking.UpdateResourceResponse{Err: fmt.Errorf("failed to commit transaction: %w", err)}
	}

	return booking.UpdateResourceResponse{
		Resource: r.toModel(),
	}
}

// DeleteResource permanently removes a resource by ID. Only the resource owner may delete a
// resource. Returns ERESOURCENOTFOUND if the resource does not exist or the user does not have
// permission to delete it.
func (s *resourceService) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) booking.DeleteResourceResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.DeleteResourceResponse{Err: fmt.Errorf("failed to start transaction: %w", err)}
	}

	err = deleteResource(ctx, tx, req.ID)
	if err != nil {
		return booking.DeleteResourceResponse{Err: fmt.Errorf("failed to delete resource")}
	}

	return booking.DeleteResourceResponse{}
}

// findResourceByID retrieves a single resource by ID from the database.
func findResourceByID(
	ctx context.Context,
	tx *Tx,
	id int,
	withEdges func(*ResourceQuery) *ResourceQuery,
) (*Resource, error) {
	q := tx.Resource.
		Query().
		Where(resource.ID(id))
	if withEdges != nil {
		withEdges(q)
	}

	r, err := q.First(ctx)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.ERESOURCENOTFOUND, "Could not find resource with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	return r, nil
}

// findResources retrieves a slice of resources from the database based on a set
// of queryable parameters. Also returns the total number of records that meet
// the critieria before offsets and limits are applied.
func findResources(
	ctx context.Context,
	tx *Tx,
	req booking.FindResourcesRequest,
	withEdges func(*ResourceQuery) *ResourceQuery,
) ([]*Resource, int, error) {
	query := tx.Resource.Query()
	if req.ID != nil {
		query.Where(resource.ID(*req.ID))
	}
	if req.Name != nil {
		query.Where(resource.Name(*req.Name))
	}
	if req.Description != nil {
		query.Where(resource.Description(*req.Description))
	}
	totalItems, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count resources: %w", err)
	}
	query = query.Offset(req.Offset)
	if req.Limit == 0 {
		query.Limit(10)
	} else {
		query.Limit(req.Limit)
	}
	if withEdges != nil {
		withEdges(query)
	}

	r, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query resources: %w", err)
	}

	return r, totalItems, nil
}

// checkForResourceNameConflict queries the database to see if a resource with a
// given name exists. Returns booking.ERESOURCENAMECONFLICT if a matching
// resource is found.
func checkForResourceNameConflict(ctx context.Context, tx *Tx, name string) error {
	count, err := tx.Resource.
		Query().
		Where(resource.Name(name)).
		Count(ctx)
	if err != nil {
		return fmt.Errorf("failed to count resources: %w", err)
	}
	if count > 0 {
		return booking.Errorf(booking.ERESOURCENAMECONFLICT, "A resource with name '%s' already exists", name)
	}
	return nil
}

// createResource creates a resource in the database. Also creates slots defined
// in the Slots parameter of the struct.
func createResource(ctx context.Context, tx *Tx, req booking.CreateResourceRequest) (*Resource, error) {
	orgID := booking.OrganizationIDFromContext(ctx)
	r, err := tx.Resource.
		Create().
		SetBookingPrice(req.BookingPrice).
		SetDescription(req.Description).
		SetName(req.Name).
		SetOrganizationID(orgID).
		SetPassword(req.Password).
		SetPrice(req.Price).
		SetTimezone(req.Timezone).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	for _, s := range req.Slots {
		ns, err := createSlot(ctx, tx, s, r)
		if err != nil {
			return nil, fmt.Errorf("failed to create slot: %w", err)
		}
		r.Edges.Slots = append(r.Edges.Slots, ns)
	}

	return r, nil
}

// createSlot creates a slot in the database.
func createSlot(ctx context.Context, tx *Tx, s *booking.Slot, r *Resource) (*Slot, error) {
	ns, err := tx.Slot.
		Create().
		SetDay(s.Day).
		SetEndTime(s.EndTime).
		SetStartTime(s.StartTime).
		SetNillableQuantity(s.Quantity).
		SetResource(r).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create slot: %w", err)
	}
	return ns, nil
}

// updateResource updates a resource in the database. Also updates the slots
// associated with the resource. It achieves this by deleting the existing slot
// records and inserting new ones.
func updateResource(ctx context.Context, tx *Tx, req booking.UpdateResourceRequest) (*Resource, error) {
	r, err := tx.Resource.
		UpdateOneID(req.ID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetTimezone(req.Timezone).
		SetPassword(req.Password).
		SetPrice(req.Price).
		SetBookingPrice(req.BookingPrice).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update resource: %w", err)
	}

	err = deleteResourceSlots(ctx, tx, r.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete resource slots: %w", err)
	}

	for _, s := range req.Slots {
		ns, err := createSlot(ctx, tx, s, r)
		if err != nil {
			return nil, fmt.Errorf("failed to create slot: %w", err)
		}
		r.Edges.Slots = append(r.Edges.Slots, ns)
	}

	return r, nil
}

// deleteResourceSlots deletes all slots associated with a given resource.
func deleteResourceSlots(ctx context.Context, tx *Tx, rid int) error {
	_, err := tx.Slot.Delete().Where(slot.ResourceId(rid)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete slots for resource: %w", err)
	}
	return nil
}

func deleteResource(ctx context.Context, tx *Tx, id int) error {
	err := tx.Resource.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}

	err = deleteResourceSlots(ctx, tx, id)
	if err != nil {
		return fmt.Errorf("failed to delete resource slots: %w", err)
	}

	return nil
}

func (r *Resource) toModel() *booking.Resource {
	result := &booking.Resource{
		ID:             r.ID,
		OrganizationID: r.OrganizationId,
		Name:           r.Name,
		Description:    r.Description,
		Timezone:       r.Timezone,
		Password:       r.Password,
		Price:          r.Price,
		BookingPrice:   r.BookingPrice,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}

	if r.Edges.Organization != nil {
		result.Organization = r.Edges.Organization.toModel()
	}

	if r.Edges.Slots != nil {
		result.Slots = make([]*booking.Slot, 0)
		for _, s := range r.Edges.Slots {
			result.Slots = append(result.Slots, s.toModel())
		}
	} else {
		result.Slots = make([]*booking.Slot, 0)
	}

	return result
}

func (r Resources) toModels() []*booking.Resource {
	var resources []*booking.Resource
	for _, v := range r {
		resources = append(resources, v.toModel())
	}
	return resources
}

func (s *Slot) toModel() *booking.Slot {
	return &booking.Slot{
		Day:       s.Day,
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		Quantity:  s.Quantity,
	}
}
