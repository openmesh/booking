package ent

import (
	"context"
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

func (s *resourceService) FindResourceByID(
	ctx context.Context,
	req booking.FindResourceByIDRequest,
) booking.FindResourceByIDResponse {
	r, err := s.client.Resource.
		Query().
		Where(resource.ID(req.ID)).
		WithSlots().
		First(ctx)
	if _, ok := err.(*NotFoundError); ok {
		return booking.FindResourceByIDResponse{
			Err: booking.Error{
				Code:   booking.ENOTFOUND,
				Detail: fmt.Sprintf("'Resource' with ID '%d' could not be found", req.ID),
				Title:  "Resource not found",
				Params: nil,
			},
		}
	}
	if err != nil {
		return booking.FindResourceByIDResponse{Err: err}
	}
	return booking.FindResourceByIDResponse{Resource: r.toModel()}
}

func (s *resourceService) FindResources(ctx context.Context, req booking.FindResourcesRequest) booking.FindResourcesResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.FindResourcesResponse{Err: err}
	}
	query := tx.Resource.Query().
		Where(resource.OrganizationId(booking.OrganizationIDFromContext(ctx)))
	if req.ID != nil {
		query = query.Where(resource.ID(*req.ID))
	}
	if req.Name != nil {
		query = query.Where(resource.Name(*req.Name))
	}
	if req.Description != nil {
		query = query.Where(resource.Description(*req.Description))
	}
	totalItems, err := query.Count(ctx)
	if err != nil {
		return booking.FindResourcesResponse{Err: err}
	}
	query = query.Offset(req.Offset)
	if req.Limit == 0 {
		query = query.Limit(10)
	} else {
		query = query.Limit(req.Limit)
	}
	resources, err := query.WithSlots().All(ctx)
	if err != nil {
		return booking.FindResourcesResponse{Err: err}
	}

	return booking.FindResourcesResponse{
		Resources:  Resources(resources).toModels(),
		TotalItems: totalItems,
	}
}

func (s *resourceService) CreateResource(ctx context.Context, req booking.CreateResourceRequest) booking.CreateResourceResponse {
	orgID := booking.OrganizationIDFromContext(ctx)

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateResourceResponse{Err: err}
	}

	// Check for existing resource with same name.
	count, err := tx.Resource.Query().Where(resource.Name(req.Name), resource.OrganizationId(orgID)).Count(ctx)
	if err != nil {
		return booking.CreateResourceResponse{Err: err}
	}
	if count > 0 {
		return booking.CreateResourceResponse{
			Err: booking.Error{
				Code:   booking.ECONFLICT,
				Detail: "One or more validation errors occurred while processing your request.",
				Title:  "Invalid request",
				Params: []booking.ValidationError{{Name: "name", Reason: fmt.Sprintf("A resource with name '%s' already exists", req.Name)}},
			},
		}
	}

	// Create resource.
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
		return booking.CreateResourceResponse{Err: tx.Rollback()}
	}

	// Create slots for resource.
	for _, s := range req.Slots {
		slot, err := tx.Slot.
			Create().
			SetDay(s.Day).
			SetEndTime(s.EndTime).
			SetStartTime(s.StartTime).
			SetNillableQuantity(s.Quantity).
			SetResource(r).
			Save(ctx)
		if err != nil {
			return booking.CreateResourceResponse{Err: tx.Rollback()}
		}
		r.Edges.Slots = append(r.Edges.Slots, slot)
	}

	if err != nil {
		return booking.CreateResourceResponse{Err: tx.Rollback()}
	}

	return booking.CreateResourceResponse{
		Resource: r.toModel(),
		Err:      tx.Commit(),
	}
}

func (s *resourceService) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) booking.UpdateResourceResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.UpdateResourceResponse{Err: err}
	}

	// TODO check for not found

	organizationID := booking.OrganizationIDFromContext(ctx)
	count, err := tx.Resource.Query().Where(resource.Name(req.Name), resource.OrganizationId(organizationID)).Count(ctx)
	if err != nil {
		return booking.UpdateResourceResponse{Err: err}
	}
	if count > 0 {
		return booking.UpdateResourceResponse{
			Err: booking.Error{
				Code:   booking.ECONFLICT,
				Detail: "One or more validation errors occurred while processing your request.",
				Title:  "Invalid request",
				Params: []booking.ValidationError{{Name: "name", Reason: fmt.Sprintf("A resource with name '%s' already exists", req.Name)}},
			},
		}
	}

	// Update resource
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
		return booking.UpdateResourceResponse{Err: tx.Rollback()}
	}

	// Confirm that resource belongs to currently authenticated organization.
	if organizationID != r.OrganizationId {
		return booking.UpdateResourceResponse{
			Err: booking.Error{
				Code:   booking.ENOTFOUND,
				Detail: fmt.Sprintf("'Resource' with ID '%d' could not be found", req.ID),
				Title:  "Resource not found",
				Params: nil,
			},
		}
	}

	// Remove existing slots
	_, err = tx.Slot.Delete().Where(slot.ResourceId(req.ID)).Exec(ctx)
	if err != nil {
		return booking.UpdateResourceResponse{Err: tx.Rollback()}
	}

	// Create new slots for resource
	for _, s := range req.Slots {
		slot, err := tx.Slot.Create().
			SetDay(s.Day).
			SetEndTime(s.EndTime).
			SetStartTime(s.StartTime).
			SetNillableQuantity(s.Quantity).
			SetResource(r).
			Save(ctx)
		if err != nil {
			return booking.UpdateResourceResponse{Err: tx.Rollback()}
		}
		r.Edges.Slots = append(r.Edges.Slots, slot)
	}

	if err != nil {
		return booking.UpdateResourceResponse{Err: tx.Rollback()}
	}

	return booking.UpdateResourceResponse{
		Resource: r.toModel(),
		Err:      tx.Commit(),
	}
}

func (s *resourceService) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) booking.DeleteResourceResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.DeleteResourceResponse{Err: err}
	}
	orgID := booking.OrganizationIDFromContext(ctx)
	deleted, err := tx.Resource.
		Delete().
		Where(resource.ID(req.ID), resource.OrganizationId(orgID)).
		Exec(ctx)
	if err != nil {
		return booking.DeleteResourceResponse{Err: err}
	}

	if deleted == 0 {
		return booking.DeleteResourceResponse{
			Err: booking.Error{
				Code:   booking.ENOTFOUND,
				Detail: fmt.Sprintf("'Resource' with ID '%d' could not be found", req.ID),
				Title:  "Resource not found",
				Params: nil,
			},
		}
	}
	return booking.DeleteResourceResponse{}
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
