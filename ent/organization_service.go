package ent

import (
	"context"
	"github.com/openmesh/booking/ent/organization"
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

func (s *organizationService) FindOrganizationByPrivateKey(ctx context.Context, key string) (*booking.Organization, error) {
	org, err := s.client.Organization.
		Query().
		Where(organization.PrivateKey(key)).
		First(ctx)
	return org.toModel(), err
}

func (s *organizationService) UpdateOrganization(ctx context.Context, upd booking.OrganizationUpdate) (*booking.Organization, error) {
	organizationID := booking.OrganizationIDFromContext(ctx)
	updateBuilder := s.client.Organization.UpdateOneID(organizationID)

	if upd.Name != nil {
		updateBuilder.SetName(*upd.Name)
	}

	org, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return org.toModel(), nil
}

func (s *organizationService) CreateOrganization(ctx context.Context, organization *booking.Organization) error {
	tx, err := s.client.Tx(ctx)

	entity, err := tx.Organization.Create().
		SetPublicKey(randStringBytes(16)).
		SetPrivateKey(randStringBytes(16)).
		SetName(organization.Name).
		Save(ctx)

	if err != nil {
		return err
	}

	organization = entity.toModel()

	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
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
