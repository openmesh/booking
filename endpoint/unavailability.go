package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// UnavailabilityEndpoints collects all the endpoints that compose a booking.UnavailabilityService.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type UnavailabilityEndpoints struct {
	FindUnavailabilityByIDEndpoint endpoint.Endpoint
	FindUnavailabilitiesEndpoint   endpoint.Endpoint
	CreateUnavailabilityEndpoint   endpoint.Endpoint
	UpdateUnavailabilityEndpoint   endpoint.Endpoint
	DeleteUnavailabilityEndpoint   endpoint.Endpoint
}

// MakeUnavailabilityEndpoints returns a UnavailabilityEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeUnavailabilityEndpoints(s booking.UnavailabilityService) UnavailabilityEndpoints {
	return UnavailabilityEndpoints{
		FindUnavailabilityByIDEndpoint: MakeFindUnavailabilityByIDEndpoint(s),
		FindUnavailabilitiesEndpoint:   MakeFindUnavailabilitiesEndpoint(s),
		CreateUnavailabilityEndpoint:   MakeCreateUnavailabilityEndpoint(s),
		UpdateUnavailabilityEndpoint:   MakeUpdateUnavailabilityEndpoint(s),
		DeleteUnavailabilityEndpoint:   MakeDeleteUnavailabilityEndpoint(s),
	}
}

// MakeFindUnavailabilityByIDEndpoint returns an endpoint via the passed service.
func MakeFindUnavailabilityByIDEndpoint(s booking.UnavailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindUnavailabilityByID(ctx, r.(booking.FindUnavailabilityByIDRequest)), nil
	}
}

// MakeFindUnavailabilitiesEndpoint returns an endpoint via the passed service.
func MakeFindUnavailabilitiesEndpoint(s booking.UnavailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindUnavailabilities(ctx, r.(booking.FindUnavailabilitiesRequest)), nil
	}
}

// MakeCreateUnavailabilityEndpoint returns an endpoint via the passed service.
func MakeCreateUnavailabilityEndpoint(s booking.UnavailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.CreateUnavailability(ctx, r.(booking.CreateUnavailabilityRequest)), nil
	}
}

// MakeUpdateUnavailabilityEndpoint returns an endpoint via the passed service.
func MakeUpdateUnavailabilityEndpoint(s booking.UnavailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.UpdateUnavailability(ctx, r.(booking.UpdateUnavailabilityRequest)), nil
	}
}

// MakeDeleteUnavailabilityEndpoint returns an endpoint via the passed service.
func MakeDeleteUnavailabilityEndpoint(s booking.UnavailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.DeleteUnavailability(ctx, r.(booking.DeleteUnavailabilityRequest)), nil
	}
}
