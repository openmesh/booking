package booking

import (
	"context"
)

type OAuthService interface {
	GetRedirectURL(ctx context.Context, req GetRedirectURLRequest) GetRedirectURLResponse
	HandleCallback(ctx context.Context, req HandleCallbackRequest) HandleCallbackResponse
}

// GetRedirectURLRequest represents a payload used by the GetRedirectURL method of a OAuthService
type GetRedirectURLRequest struct {
	Source string `json:"source" source:"url"`
}

// Validate a GetRedirectURL. Returns a ValidationError for each requirement that fails.
func (r GetRedirectURLRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetRedirectURLResponse represents a response returned by the GetRedirectURL method of a OAuthService.
type GetRedirectURLResponse struct {
	URL   string `json:"url"`
	State string `json:"state"`
	Err   error  `json:"error,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetRedirectURLResponse) Error() error { return r.Err }

// HandleCallbackRequest represents a payload used by the HandleCallback method of a OAuthService
type HandleCallbackRequest struct {
	Source      string `json:"source"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirectUrl`
}

// Validate a HandleCallback. Returns a ValidationError for each requirement that fails.
func (r HandleCallbackRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// HandleCallbackResponse represents a response returned by the HandleCallback method of a OAuthService.
type HandleCallbackResponse struct {
	UserID      int    `json:"userId"`
	RedirectURL string `json:"redirectUrl`
	Err         error  `json:"error,omitempty`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r HandleCallbackResponse) Error() error { return r.Err }

// OAuthServiceMiddleware defines a middleware for OAuthService
type OAuthServiceMiddleware func(service OAuthService) OAuthService

// OAuthValidationMiddleware returns a middleware for validating requests made to a OAuthService
func OAuthValidationMiddleware() OAuthServiceMiddleware {
	return func(next OAuthService) OAuthService {
		return oauthValidationMiddleware{next}
	}
}

type oauthValidationMiddleware struct {
	OAuthService
}

// GetRedirectURL validates a GetRedirectURLRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw oauthValidationMiddleware) GetRedirectURL(ctx context.Context, req GetRedirectURLRequest) GetRedirectURLResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetRedirectURLResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.OAuthService.GetRedirectURL(ctx, req)
}

// HandleCallback validates a HandleCallbackRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw oauthValidationMiddleware) HandleCallback(ctx context.Context, req HandleCallbackRequest) HandleCallbackResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return HandleCallbackResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.OAuthService.HandleCallback(ctx, req)
}
