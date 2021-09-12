package ent

import (
	"context"

	"github.com/openmesh/booking"
)

type handler struct {
	client *Client
}

func NewHandler(client *Client) *handler {
	return &handler{
		client,
	}
}

func (h *handler) Handle(ctx context.Context, query booking.FindResourceByIDQuery) (*booking.Resource, error) {
	r, err := h.client.Resource.Get(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return r.toModel(), nil
}

type resourceService struct {
	client *Client
}

func NewResourceService(client *Client) *resourceService {
	return &resourceService{
		client,
	}
}

func (s *resourceService) FindResourceByID(ctx context.Context, id int) (*booking.Resource, error) {
	r, err := s.client.Resource.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.toModel(), nil
}

func (s *resourceService) FindResources(ctx context.Context, filter booking.ResourceFilter) ([]*booking.Resource, int, error) {
	return []*booking.Resource{}, 0, nil
}

func (s *resourceService) CreateResource(ctx context.Context, resource *booking.Resource) (*booking.Resource, error) {
	organizationID := booking.OrganizationIDFromContext(ctx)
	if organizationID == 0 {
		return nil,booking.Errorf(booking.EUNAUTHORIZED, "Failed to create resource: No user authenticated")
	}

	entity, err := s.client.Resource.Create().
		SetBookingPrice(resource.BookingPrice).
		SetDescription(resource.Description).
		SetName(resource.Name).
		SetOrganizationID(organizationID).
		SetPassword(resource.Password).
		SetPrice(resource.Price).
		SetTimezone(resource.Timezone).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	resource = entity.toModel()

	return resource, nil
}

func (s *resourceService) UpdateResource(ctx context.Context, id int, upd booking.ResourceUpdate) (*booking.Resource, error) {
	resource, err := s.client.Resource.
		UpdateOneID(id).
		SetName(upd.Name).
		SetDescription(upd.Description).
		SetTimezone(upd.Timezone).
		SetPassword(upd.Password).
		SetPrice(upd.Price).
		SetBookingPrice(upd.BookingPrice).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return resource.toModel(), nil
}

func (s *resourceService) DeleteResource(ctx context.Context, id int) error {
	return s.client.Resource.DeleteOneID(id).Exec(ctx)
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
		for _, s := range r.Edges.Slots {
			result.Slots = append(result.Slots, s.toModel())
		}
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
