package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// ResourceEndpoints collects all the endpoints that compose a resource service.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type ResourceEndpoints struct {
	FindResourceByIDEndpoint endpoint.Endpoint
	FindResourcesEndpoint    endpoint.Endpoint
	CreateResourceEndpoint   endpoint.Endpoint
	UpdateResourceEndpoint   endpoint.Endpoint
	DeleteResourceEndpoint   endpoint.Endpoint
}

// MakeResourceEndpoints returns a ResourceEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeResourceEndpoints(s booking.ResourceService) ResourceEndpoints {
	return ResourceEndpoints{
		FindResourceByIDEndpoint: MakeFindResourceByIDEndpoint(s),
		FindResourcesEndpoint:    MakeFindResourcesEndpoint(s),
		CreateResourceEndpoint:   MakeCreateResourceEndpoint(s),
		UpdateResourceEndpoint:   MakeUpdateResourceEndpoint(s),
		DeleteResourceEndpoint:   MakeDeleteResourceEndpoint(s),
	}
}

// MakeFindResourceByIDEndpoint returns an endpoint via the passed service.
func MakeFindResourceByIDEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindResourceByID(ctx, r.(booking.FindResourceByIDRequest)), nil
	}
}

// MakeFindResourcesEndpoint returns an endpoint via the passed service.
func MakeFindResourcesEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindResources(ctx, r.(booking.FindResourcesRequest)), nil
	}
}

// MakeCreateResourceEndpoint returns an endpoint via the passed service.
func MakeCreateResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.CreateResource(ctx, r.(booking.CreateResourceRequest)), nil
	}
}

// MakeUpdateResourceEndpoint returns an endpoint via the passed service.
func MakeUpdateResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.UpdateResource(ctx, r.(booking.UpdateResourceRequest)), nil
	}
}

// MakeDeleteResourceEndpoint returns an endpoint via the passed service.
func MakeDeleteResourceEndpoint(s booking.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.DeleteResource(ctx, r.(booking.DeleteResourceRequest)), nil
	}
}
