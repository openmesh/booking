package ent

import (
	"context"
	"math/rand"

	"github.com/openmesh/booking"
)

type organizationService struct {
	client *Client
}

func NewOrganizationService(client *Client) *organizationService {
	return &organizationService{
		client,
	}
}

func (s *organizationService) FindCurrentOrganization(ctx context.Context) (*booking.Organization, error) {
	organizationID := booking.OrganizationIDFromContext(ctx)
	o, err := s.client.Organization.Get(ctx, organizationID)
	if err != nil {
		return nil, err
	}
	return o.toModel(), nil
}

func (s *organizationService) UpdateOrganization(ctx context.Context, upd booking.OrganizationUpdate) (*booking.Organization, error) {
	organizationID := booking.OrganizationIDFromContext(ctx)
	updateBuilder := s.client.Organization.UpdateOneID(organizationID)

	if upd.Name != nil {
		updateBuilder.SetName(*upd.Name)
	}
	if upd.OwnerID != nil {
		updateBuilder.SetOwnerID(*upd.OwnerID)
	}

	organization, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return organization.toModel(), nil
}

func (s *organizationService) CreateOrganization(ctx context.Context, organization *booking.Organization) error {
	entity, err := s.client.Organization.Create().
		SetPublicKey(generateKey()).
		// SetPrivateKey(generateKey()).
		SetName(organization.Name).
		SetOwnerID(organization.OwnerID).
		Save(ctx)

	if err != nil {
		return err
	}

	organization = entity.toModel()

	return nil
}

func generateKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return string(b)
}

func (o *Organization) toModel() *booking.Organization {
	return &booking.Organization{
		ID:        o.ID,
		Name:      o.Name,
		PublicKey: o.PublicKey,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
