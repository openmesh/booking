package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// TokenEndpoints collects all the endpoints that compose a booking.TokenService.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type TokenEndpoints struct {
	CreateTokenEndpoint endpoint.Endpoint
	FindTokensEndpoint  endpoint.Endpoint
}

// MakeTokenEndpoints returns a TokenEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeTokenEndpoints(s booking.TokenService) TokenEndpoints {
	return TokenEndpoints{
		CreateTokenEndpoint: MakeCreateTokenEndpoint(s),
		FindTokensEndpoint:  MakeFindTokensEndpoint(s),
	}
}

// MakeCreateTokenEndpoint returns an endpoint via the passed service.
func MakeCreateTokenEndpoint(s booking.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.CreateToken(ctx, r.(booking.CreateTokenRequest)), nil
	}
}

// MakeFindTokensEndpoint returns an endpoint via the passed service.
func MakeFindTokensEndpoint(s booking.TokenService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.FindTokens(ctx, r.(booking.FindTokensRequest)), nil
	}
}
