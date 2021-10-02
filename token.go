package booking

import (
	"context"
	"time"
)

type Token struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Expiry         *time.Time `json:"expiry"`
	UserID         int        `json:"userId"`
	OrganizationID int        `json:"organizationId"`
}

type TokenService interface {
	CreateToken(ctx context.Context, req CreateTokenRequest) CreateTokenResponse
	FindTokens(ctx context.Context, req FindTokensRequest) FindTokensResponse
}

// CreateTokenRequest represents a payload used by the CreateToken method of a TokenService
type CreateTokenRequest struct {
	Name   string     `json:"name"`
	Expiry *time.Time `json:"expiry"`
}

// Validate a CreateToken. Returns a ValidationError for each requirement that fails.
func (r CreateTokenRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// CreateTokenResponse represents a response returned by the CreateToken method of a TokenService.
type CreateTokenResponse struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Expiry         *time.Time `json:"expiry"`
	UserID         int        `json:"userId"`
	OrganizationId int        `json:"organizationId"`
	Err            error      `json:"err"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateTokenResponse) Error() error { return r.Err }

// FindTokensRequest represents a payload used by the FindTokens method of a TokenService
type FindTokensRequest struct{}

// Validate a FindTokens. Returns a ValidationError for each requirement that fails.
func (r FindTokensRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindTokensResponse represents a response returned by the FindTokens method of a TokenService.
type FindTokensResponse struct {
	Tokens     []Token `json:"tokens,omitempty"`
	TotalItems int     `json:"totalItems"`
	Err        error   `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindTokensResponse) Error() error { return r.Err }

// TokenServiceMiddleware defines a middleware for TokenService
type TokenServiceMiddleware func(service TokenService) TokenService

// TokenValidationMiddleware returns a middleware for validating requests made to a TokenService
func TokenValidationMiddleware() TokenServiceMiddleware {
	return func(next TokenService) TokenService {
		return tokenValidationMiddleware{next}
	}
}

type tokenValidationMiddleware struct {
	TokenService
}

// CreateToken validates a CreateTokenRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw tokenValidationMiddleware) CreateToken(ctx context.Context, req CreateTokenRequest) CreateTokenResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateTokenResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.TokenService.CreateToken(ctx, req)
}

// FindTokens validates a FindTokensRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw tokenValidationMiddleware) FindTokens(ctx context.Context, req FindTokensRequest) FindTokensResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindTokensResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.TokenService.FindTokens(ctx, req)
}
