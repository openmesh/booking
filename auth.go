package booking

import (
	"context"
	"fmt"
	"time"
)

// Authentication providers. Currently we only support GitHub but any OAuth
// provider could be supported.
const (
	// AuthSourceGitHub is the key used to represent the GitHub authentication source.
	AuthSourceGitHub = "github"
)

// Auth represents a set of OAuth credentials. These are linked to a User so a
// single user could authenticate through multiple providers.
//
// The authentication system links users by email address, however, some GitHub
// users don't provide their email publicly so we may not be able to link them
// by email address. It's a moot point, however, as we only support GitHub as an
// OAuth provider.
type Auth struct {
	ID int `json:"id"`

	// User can have one or more methods of authentication. However, only one per
	// source is allowed per user.
	UserID int   `json:"userId"`
	User   *User `json:"user"`

	// The authentication source & the source provider's user ID. Source can only
	// be "github" currently.
	Source   string `json:"source"`
	SourceID string `json:"sourceId"`

	// OAuth fields returned from the authentication provider. GitHub does not use
	// refresh tokens but the field exists for future providers.
	AccessToken  string     `json:"-"`
	RefreshToken string     `json:"-"`
	Expiry       *time.Time `json:"-"`

	// Timestamps of creation & last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AvatarURL returns a URL to the avatar image hosted by the authentication
// source. Returns an empty string if the authentication source is invalid.
func (a *Auth) AvatarURL(size int) string {
	switch a.Source {
	case AuthSourceGitHub:
		return fmt.Sprintf("https://avatars1.githubusercontent.com/u/%s?s=%d", a.SourceID, size)
	default:
		return ""
	}
}

type AuthService interface {
	// FindAuthByID looks up an authentication object by ID along with the associated user.
	// Returns ENOTFOUND if ID does not exist.
	FindAuthByID(ctx context.Context, req FindAuthByIDRequest) FindAuthByIDResponse
	// FindAuths retrieves authentication objects based on a filter. Also returns the total
	// number of objects that match the filter. This may differ from the returned
	// object count if the Limit field is set.
	FindAuths(ctx context.Context, req FindAuthsRequest) FindAuthsResponse
	// CreateAuth creates a new authentication object If a User is attached to auth, then the
	// auth object is linked to an existing user. Otherwise a new user object is
	// created.
	//
	// Returns the created Auth.
	CreateAuth(ctx context.Context, req CreateAuthRequest) CreateAuthResponse
	// UpdateAuth updates the refresh token value for an auth.
	UpdateAuth(ctx context.Context, req UpdateAuthRequest) UpdateAuthResponse
	// DeleteAuth permanently deletes an authentication object from the system by ID. The
	// parent user object is not removed.
	DeleteAuth(ctx context.Context, req DeleteAuthRequest) DeleteAuthResponse
}

// FindAuthByIDRequest represents a payload used by the FindAuthByID method of a AuthService
type FindAuthByIDRequest struct {
	ID int `json:"id"`
}

// Validate a FindAuthByID. Returns a ValidationError for each requirement that fails.
func (r FindAuthByIDRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{
			{
				Name:   "id",
				Reason: "Must be at least 1",
			},
		}
	}
	return nil
}

// FindAuthByIDResponse represents a response returned by the FindAuthByID method of a AuthService.
type FindAuthByIDResponse struct {
	*Auth
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindAuthByIDResponse) Error() error {
	return r.Err
}

// FindAuthsRequest represents a payload used by the FindAuths method of a AuthService
type FindAuthsRequest struct {
	// Filtering fields.
	ID       *int    `json:"id"`
	UserID   *int    `json:"userId"`
	Source   *string `json:"source"`
	SourceID *string `json:"sourceId"`

	// Restricts results to a subset of the total range. Can be used for
	// pagination.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// Validate a FindAuths. Returns a ValidationError for each requirement that fails.
func (r FindAuthsRequest) Validate() []ValidationError {
	errs := make([]ValidationError, 0)
	if r.ID != nil && *r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.UserID != nil && *r.UserID < 1 {
		errs = append(errs, ValidationError{Name: "userId", Reason: "Must be at least 1"})
	}
	if r.Offset < 0 {
		errs = append(errs, ValidationError{Name: "offset", Reason: "Must be greater than or equal to 0"})
	}
	if r.Limit < 0 {
		errs = append(errs, ValidationError{Name: "limit", Reason: "Must be greater than or equal to 0"})
	}
	validSources := []string{AuthSourceGitHub}
	if r.Source != nil && !Strings(validSources).contains(*r.Source) {
		errs = append(errs, ValidationError{Name: "source", Reason: "Must be a valid auth source"})
	}

	return errs
}

// FindAuthsResponse represents a response returned by the FindAuths method of a AuthService.
type FindAuthsResponse struct {
	Auths      []*Auth `json:"resources,omitempty"`
	TotalItems int     `json:"totalItems"`
	Err        error   `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindAuthsResponse) Error() error { return r.Err }

// CreateAuthRequest represents a payload used by the CreateAuth method of a AuthService
type CreateAuthRequest struct {
	UserID       int        `json:"userId"`
	UserName     string     `json:"userName"`
	UserEmail    string     `json:"userEmail"`
	Source       string     `json:"source"`
	SourceID     string     `json:"sourceId"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
	Expiry       *time.Time `json:"expiry"`
}

// Validate a CreateAuth. Returns a ValidationError for each requirement that fails.
func (r CreateAuthRequest) Validate() []ValidationError {
	// insert validation logic here
	errs := make([]ValidationError, 0)
	if r.UserEmail == "" {
		errs = append(errs, ValidationError{Name: "userEmail", Reason: "User email required"})
	}
	if r.Source == "" {
		errs = append(errs, ValidationError{Name: "source", Reason: "Source required"})
	}
	if r.SourceID == "" {
		errs = append(errs, ValidationError{Name: "sourceId", Reason: "Source ID required"})
	}
	if r.AccessToken == "" {
		errs = append(errs, ValidationError{Name: "accessToken", Reason: "Access token required"})
	}
	return errs
}

// CreateAuthResponse represents a response returned by the CreateAuth method of a AuthService.
type CreateAuthResponse struct {
	*Auth
	Err error
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateAuthResponse) Error() error { return r.Err }

// UpdateAuthRequest represents a payload used by the UpdateRefreshToken method of a AuthService
type UpdateAuthRequest struct {
	ID           int        `json:"id"`
	RefreshToken string     `json:"refreshToken"`
	AccessToken  string     `json:"accessToken"`
	Expiry       *time.Time `json:"expiry"`
}

// Validate a UpdateRefreshToken. Returns a ValidationError for each requirement that fails.
func (r UpdateAuthRequest) Validate() []ValidationError {
	var errs []ValidationError
	if r.ID < 1 {
		errs = append(errs, ValidationError{Name: "id", Reason: "Must be at least 1"})
	}
	if r.RefreshToken == "" {
		errs = append(errs, ValidationError{Name: "name", Reason: "Name is required"})
	}
	return errs
}

// UpdateAuthResponse represents a response returned by the UpdateRefreshToken method of a AuthService.
type UpdateAuthResponse struct {
	*Auth `json:"auth,omitempty"`
	Err   error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateAuthResponse) Error() error { return r.Err }

// DeleteAuthRequest represents a payload used by the DeleteAuth method of a AuthService
type DeleteAuthRequest struct {
	ID int `json:"id"`
}

// Validate a DeleteAuth. Returns a ValidationError for each requirement that fails.
func (r DeleteAuthRequest) Validate() []ValidationError {
	if r.ID < 1 {
		return []ValidationError{{Name: "id", Reason: "Must be at least 1"}}
	}
	return nil
}

// DeleteAuthResponse represents a response returned by the DeleteAuth method of a AuthService.
type DeleteAuthResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteAuthResponse) Error() error { return r.Err }

// AuthServiceMiddleware defines a middleware for AuthService
type AuthServiceMiddleware func(service AuthService) AuthService

// AuthValidationMiddleware returns a middleware for validating requests made to a AuthService
func AuthValidationMiddleware() AuthServiceMiddleware {
	return func(next AuthService) AuthService {
		return authValidationMiddleware{next}
	}
}

type authValidationMiddleware struct {
	AuthService
}

// FindAuthByID validates a FindAuthByIDRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw authValidationMiddleware) FindAuthByID(ctx context.Context, req FindAuthByIDRequest) FindAuthByIDResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindAuthByIDResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.AuthService.FindAuthByID(ctx, req)
}

// FindAuths validates a FindAuthsRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw authValidationMiddleware) FindAuths(ctx context.Context, req FindAuthsRequest) FindAuthsResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindAuthsResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.AuthService.FindAuths(ctx, req)
}

// CreateAuth validates a CreateAuthRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw authValidationMiddleware) CreateAuth(ctx context.Context, req CreateAuthRequest) CreateAuthResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateAuthResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.AuthService.CreateAuth(ctx, req)
}

// UpdateRefreshToken validates a UpdateRefreshTokenRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw authValidationMiddleware) UpdateRefreshToken(ctx context.Context, req UpdateAuthRequest) UpdateAuthResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return UpdateAuthResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.AuthService.UpdateAuth(ctx, req)
}

// DeleteAuth validates a DeleteAuthRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw authValidationMiddleware) DeleteAuth(ctx context.Context, req DeleteAuthRequest) DeleteAuthResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return DeleteAuthResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.AuthService.DeleteAuth(ctx, req)
}
