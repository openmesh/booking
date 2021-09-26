package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// OAuthEndpoints collects all the endpoints that compose a booking.OAuthService.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type OAuthEndpoints struct {
	GetRedirectURLEndpoint endpoint.Endpoint
	HandleCallbackEndpoint endpoint.Endpoint
}

// MakeOAuthEndpoints returns a OAuthEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeOAuthEndpoints(s booking.OAuthService) OAuthEndpoints {
	return OAuthEndpoints{
		GetRedirectURLEndpoint: MakeGetRedirectURLEndpoint(s),
		HandleCallbackEndpoint: MakeHandleCallbackEndpoint(s),
	}
}

// MakeGetRedirectURLEndpoint returns an endpoint via the passed service.
func MakeGetRedirectURLEndpoint(s booking.OAuthService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetRedirectURL(ctx, r.(booking.GetRedirectURLRequest)), nil
	}
}

// MakeHandleCallbackEndpoint returns an endpoint via the passed service.
func MakeHandleCallbackEndpoint(s booking.OAuthService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.HandleCallback(ctx, r.(booking.HandleCallbackRequest)), nil
	}
}
