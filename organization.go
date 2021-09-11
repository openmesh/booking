package booking

import (
	"context"
	"time"
)

type Organization struct {
	ID int `json:"id"`

	// Organization name.
	Name string `json:"name"`

	// A public key used to identify the organization from client facing
	// applications.
	PublicKey string `json:"publicKey"`

	// A private key used to identify the organization from server applicationvs.
	PrivateKey string `json:"privateKey"`

	// Owner of the organization. This is the user that created the organization
	// by default. However ownership can be transferred to other members of the
	// organization.
	OwnerID int   `json:"ownerId"`
	Owner   *User `json:"owner"`

	// Timestamps for user creation & last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// OrganizationService represents a service for managing organizations.
type OrganizationService interface {
	// Retrieves the organization of the currently authenticated user.
	FindCurrentOrganization(ctx context.Context) (*Organization, error)

	// Creates a new organization.
	CreateOrganization(ctx context.Context, organization *Organization) error

	// Updates the organization associated with the currently authenicated user.
	UpdateOrganization(ctx context.Context, upd OrganizationUpdate) (*Organization, error)
}

// OrganizationUpdate represents a set of fields to update on an organization.
type OrganizationUpdate struct {
	Name    *string `json:"name"`
	OwnerID *int    `json:"ownerId"`
}

// OrganizationServiceMiddleware defines a middleware for an organization service.
type OrganizationServiceMiddleware func(OrganizationService) OrganizationService