package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// OrganizationEndpoints collects all the endpoints that compose an
// organization service. It's used as a helper struct, to collect all the
// endpoints into a single parameter.
type OrganizationEndpoints struct {
	FindCurrentOrganizationEndpoint      endpoint.Endpoint
	FindOrganizationByPrivateKeyEndpoint endpoint.Endpoint
	CreateOrganizationEndpoint           endpoint.Endpoint
	UpdateOrganizationEndpoint           endpoint.Endpoint
}

// MakeOrganizationEndpoints returns an OrganizationEndpoints struct where
// each endpoint invokes the corresponding method on the provided service.
func MakeOrganizationEndpoints(s booking.OrganizationService) OrganizationEndpoints {
	return OrganizationEndpoints{
		FindCurrentOrganizationEndpoint: MakeFindCurrentOrganizationEndpoint(s),
		CreateOrganizationEndpoint:      MakeCreateOrganizationEndpoint(s),
		UpdateOrganizationEndpoint:      MakeUpdateOrganizationEndpoint(s),
	}
}

// MakeFindCurrentOrganizationEndpoint returns an endpoint via the passed service.
func MakeFindCurrentOrganizationEndpoint(s booking.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		organization, err := s.FindCurrentOrganization(ctx)
		return FindCurrentOrganizationResponse{Organization: organization, Err: err}, nil
	}
}

// MakeCreateOrganizationEndpoint returns an endpoint via the passed service.
func MakeCreateOrganizationEndpoint(s booking.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrganizationRequest)
		err := s.CreateOrganization(ctx, req.Organization)
		return CreateOrganizationResponse{Organization: req.Organization, Err: err}, nil
	}
}

// MakeUpdateOrganizationEndpoint returns an endpoint via the passed service.
func MakeUpdateOrganizationEndpoint(s booking.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateOrganizationRequest)
		organization, err := s.UpdateOrganization(ctx, req.Update)
		return UpdateOrganizationResponse{Organization: organization, Err: err}, nil
	}
}

type FindCurrentOrganizationResponse struct {
	Organization *booking.Organization `json:"organization,omitempty"`
	Err          error                 `json:"err,omitempty"`
}

type FindOrganizationByPrivateKeyRequest struct {
	Key string
}

type CreateOrganizationRequest struct {
	Organization *booking.Organization
}

type CreateOrganizationResponse struct {
	Organization *booking.Organization `json:"organization,omitempty"`
	Err          error                 `json:"err,omitempty"`
}

type UpdateOrganizationRequest struct {
	Update booking.OrganizationUpdate
}

type UpdateOrganizationResponse struct {
	Organization *booking.Organization `json:"organization,omitempty"`
	Err          error                 `json:"err,omitempty"`
}
