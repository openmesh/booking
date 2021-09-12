package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// ResourceEndpoints collects all the endpoints that compose a
// resource service. It's used as a helper struct, to collect all the
// endpoints into a single parameter.
type ResourceEndpoints struct {
	FindResourceByIDEndpoint endpoint.Endpoint
	FindResourcesEndpoint    endpoint.Endpoint
	CreateResourceEndpoint   endpoint.Endpoint
	UpdateResourceEndpoint   endpoint.Endpoint
	DeleteResourceEndpoint   endpoint.Endpoint
}

// MakeResourceEndpoints returns a ResourceEndpoints struct where
// each endpoint invokes the corresponding method on the provided service.
func MakeResourceEndpoints(s booking.ResourceService) ResourceEndpoints {
	return ResourceEndpoints{
		FindResourcesEndpoint:    MakeFindResourcesEndpoint(s),
		CreateResourceEndpoint:   MakeCreateResourceEndpoint(s),
		UpdateResourceEndpoint:   MakeUpdateResourceEndpoint(s),
		DeleteResourceEndpoint:   MakeDeleteResourceEndpoint(s),
	}
}

// MakeFindResourceByIDEndpoint returns an endpoint via the passed service.
func MakeFindResourceByIDEndpoint(h booking.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(booking.FindResourceByIDQuery)
		resource, err := h.Handle(ctx, req)
		return FindResourceByIDResponse{Resource: resource, Err: err}, nil
	}
}

// MakeFindResourcesEndpoint returns an endpoint via the passed service.
func MakeFindResourcesEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindResourcesRequest)
		resources, totalItems, err := s.FindResources(ctx, req.Filter)
		return FindResourcesResponse{Resources: resources, TotalItems: totalItems, Err: err}, nil
	}
}

// MakeCreateResourceEndpoint returns an endpoint via the passed service.
func MakeCreateResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateResourceRequest)
		resource, err := s.CreateResource(ctx, req.Resource)
		return CreateResourceResponse{Resource: resource, Err: err}, nil
	}
}

// MakeUpdateResourceEndpoint returns an endpoint via the passed service.
func MakeUpdateResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateResourceRequest)
		resource, err := s.UpdateResource(ctx, req.ID, req.Update)
		return UpdateResourceResponse{Resource: resource, Err: err}, nil
	}
}

// MakeDeleteResourceEndpoint returns an endpoint via the passed service.
func MakeDeleteResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteResourceRequest)
		err := s.DeleteResource(ctx, req.ID)
		return DeleteResourceResponse{Err: err}, nil
	}
}

type FindResourceByIDRequest struct {
	ID int
}

type FindResourceByIDResponse struct {
	Resource *booking.Resource `json:"resource,omitempty"`
	Err      error             `json:"err,omitempty"`
}

type FindResourcesRequest struct {
	Filter booking.ResourceFilter
}

type FindResourcesResponse struct {
	Resources  []*booking.Resource `json:"resources,omitempty"`
	TotalItems int                 `json:"totalItems"`
	Err        error               `json:"err,omitempty"`
}

type CreateResourceRequest struct {
	Resource *booking.Resource
}

type CreateResourceResponse struct {
	Resource *booking.Resource `json:"resource,omitempty"`
	Err      error             `json:"err,omitempty"`
}

type UpdateResourceRequest struct {
	ID     int
	Update booking.ResourceUpdate
}

type UpdateResourceResponse struct {
	Resource *booking.Resource `json:"resource,omitempty"`
	Err      error             `json:"err,omitempty"`
}

type DeleteResourceRequest struct {
	ID int
}

type DeleteResourceResponse struct {
	Err error `json:"err,omitempty"`
}
