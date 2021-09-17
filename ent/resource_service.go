package ent

import (
	"context"
	"fmt"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent/resource"
)

type resourceService struct {
	client *Client
}

func NewResourceService(client *Client) *resourceService {
	return &resourceService{
		client,
	}
}

func (s *resourceService) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) (*booking.Resource, error) {
	r, err := s.client.Resource.
		Query().
		Where(resource.ID(req.ID)).
		WithSlots().
		First(ctx) // Get(ctx, req.ID)
	if _, ok := err.(*NotFoundError); ok {
		return nil, booking.Error{
			Code:   booking.ENOTFOUND,
			Detail: fmt.Sprintf("'Resource' with ID '%d' could not be found", req.ID),
			Title:  "Resource not found",
			Params: nil,
		}
	}
	if err != nil {
		return nil, err
	}
	return r.toModel(), nil
}

func (s *resourceService) FindResources(ctx context.Context, req booking.FindResourcesRequest) ([]*booking.Resource, int, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, 0, err
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
		return nil, 0, err
	}
	query.Offset(req.Offset)
	if req.Limit == 0 {
		query.Limit(10)
	} else {
		query.Limit(req.Limit)
	}
	resources, err := query.WithSlots().All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return Resources(resources).toModels(), totalItems, nil
}

func (s *resourceService) CreateResource(ctx context.Context, req booking.CreateResourceRequest) (*booking.Resource, error) {
	organizationID := booking.OrganizationIDFromContext(ctx)
	if organizationID == 0 {
		return nil, booking.Errorf(booking.EUNAUTHORIZED, "Failed to create resource: No user authenticated")
	}
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create resource
	r, err := tx.Resource.Create().
		SetBookingPrice(req.BookingPrice).
		SetDescription(req.Description).
		SetName(req.Name).
		SetOrganizationID(organizationID).
		SetPassword(req.Password).
		SetPrice(req.Price).
		SetTimezone(req.Timezone).
		Save(ctx)

	// Create slots for resource
	for _, s := range req.Slots {
		slot, err := tx.Slot.Create().
			SetDay(s.Day).
			SetEndTime(s.EndTime).
			SetStartTime(s.StartTime).
			SetNillableQuantity(s.Quantity).
			SetResource(r).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		r.Edges.Slots = append(r.Edges.Slots, slot)
	}

	if err != nil {
		return nil, err
	}

	res := r.toModel()

	return res, nil
}

func (s *resourceService) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) (*booking.Resource, error) {
	res, err := s.client.Resource.
		UpdateOneID(req.ID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetTimezone(req.Timezone).
		SetPassword(req.Password).
		SetPrice(req.Price).
		SetBookingPrice(req.BookingPrice).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return res.toModel(), nil
}

func (s *resourceService) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) error {
	return s.client.Resource.DeleteOneID(req.ID).Exec(ctx)
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

func (s *Slot) toModel() *booking.Slot {
	return &booking.Slot{
		Day:       s.Day,
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		Quantity:  s.Quantity,
	}
}

func (r Resources) toModels() []*booking.Resource {
	var resources []*booking.Resource
	for _, v := range r {
		resources = append(resources, v.toModel())
	}
	return resources
}
