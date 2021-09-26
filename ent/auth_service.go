package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent/auth"
	"github.com/openmesh/booking/ent/user"
)

type authService struct {
	client *Client
}

// NewAuthService constructs a new instance of a booking.AuthService using ent
// as its persistence layer.
func NewAuthService(client *Client) *authService {
	return &authService{
		client,
	}
}

// FindAuthByID looks up an authentication object by ID along with the associated user.
// Returns ENOTFOUND if ID does not exist.
func (s *authService) FindAuthByID(ctx context.Context, req booking.FindAuthByIDRequest) booking.FindAuthByIDResponse {
	panic("not implemented") // TODO: Implement
}

// FindAuths retrieves authentication objects based on a filter. Also returns the total
// number of objects that match the filter. This may differ from the returned
// object count if the Limit field is set.
func (s *authService) FindAuths(ctx context.Context, req booking.FindAuthsRequest) booking.FindAuthsResponse {
	panic("not implemented") // TODO: Implement
}

// CreateAuth creates a new authentication object If a User is attached to auth, then the
// auth object is linked to an existing user. Otherwise a new user object is
// created.
//
// Returns the created Auth.
func (s *authService) CreateAuth(ctx context.Context, req booking.CreateAuthRequest) booking.CreateAuthResponse {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return booking.CreateAuthResponse{
			Err: fmt.Errorf("failed to start transaction: %w", err),
		}
	}
	defer tx.Rollback()

	// Check to see if the auth exists for the given source.
	other, err := findAuthBySourceID(ctx, tx, req.Source, req.SourceID)
	if err == nil {
		other, err = updateAuth(
			ctx,
			tx,
			booking.UpdateAuthRequest{
				ID:           other.ID,
				RefreshToken: req.RefreshToken,
				AccessToken:  req.AccessToken,
				Expiry:       req.Expiry,
			},
			func(a *Auth) (*Auth, error) {
				a.Edges.User, err = a.QueryUser().First(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed to query user: %w", err)
				}
				return a, nil
			},
		)
		if err != nil {
			return booking.CreateAuthResponse{
				Err: fmt.Errorf("cannot update auth: id=%d err=%w", other.ID, err),
			}
		}
		err = tx.Commit()
		if err != nil {
			return booking.CreateAuthResponse{
				Err: fmt.Errorf("failed to commit transaction: %w", err),
			}
		}
		return booking.CreateAuthResponse{
			Auth: other.toModel(),
		}
	} else if booking.ErrorCode(err) != booking.EAUTHNOTFOUND {
		return booking.CreateAuthResponse{
			Err: fmt.Errorf("failed to find auth by source id: %w", err),
		}
	}

	var user *User

	// Check if auth is for a new user. It is considered "new" if the caller
	// doesn't have a value for UserID.
	if req.UserID == 0 {
		// Look up the user by email. If no user can be found then create a new user
		// for the auth.
		user, err = findUserByEmail(ctx, tx, req.UserEmail)
		if err != nil {
			if booking.ErrorCode(err) == booking.EUSERNOTFOUND {
				user, err = createUser(ctx, tx, req.UserName, req.UserEmail)
				if err != nil {
					return booking.CreateAuthResponse{
						Err: fmt.Errorf("failed to create user: %w", err),
					}
				}
			} else {
				return booking.CreateAuthResponse{
					Err: fmt.Errorf("failed to find user by email: %w", err),
				}
			}
		}
	} else {
		user, err = findUserByID(ctx, tx, req.UserID)
		if err != nil {
			return booking.CreateAuthResponse{
				Err: fmt.Errorf("failed to find user by id: %w", err),
			}
		}
	}

	req.UserID = user.ID
	auth, err := createAuth(ctx, tx, req, nil)
	auth.Edges.User = user

	err = tx.Commit()
	if err != nil {
		return booking.CreateAuthResponse{
			Err: fmt.Errorf("failed to create auth: %w", err),
		}
	}

	return booking.CreateAuthResponse{
		Auth: auth.toModel(),
	}
}

func createAuth(
	ctx context.Context,
	tx *Tx,
	req booking.CreateAuthRequest,
	attachEdges func(*Auth) (*Auth, error),
) (*Auth, error) {
	a, err := tx.Auth.
		Create().
		SetAccessToken(req.AccessToken).
		SetRefreshToken(req.RefreshToken).
		SetNillableExpiry(req.Expiry).
		SetSource(req.Source).
		SetSourceId(req.SourceID).
		SetUserID(req.UserID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	if attachEdges != nil {
		a, err = attachEdges(a)
		if err != nil {
			return nil, fmt.Errorf("failed to attach edges to auth: %w", err)
		}
	}

	return a, nil
}

func findUserByID(ctx context.Context, tx *Tx, id int) (*User, error) {
	u, err := tx.User.Get(ctx, id)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.EUSERNOTFOUND, "Could not find user with ID %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return u, nil
}

func createUser(ctx context.Context, tx *Tx, name, email string) (*User, error) {
	u, err := tx.User.
		Create().
		SetEmail(email).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return u, nil
}

func findUserByEmail(ctx context.Context, tx *Tx, email string) (*User, error) {
	u, err := tx.User.
		Query().
		Where(user.Email(email)).
		First(ctx)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.EUSERNOTFOUND, "Could not find user by email: email=%s, err=%w", email, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return u, nil
}

func findAuthBySourceID(ctx context.Context, tx *Tx, source, sourceID string) (*Auth, error) {
	a, err := tx.Auth.
		Query().
		Where(
			auth.Source(source),
			auth.SourceId(sourceID),
		).
		First(ctx)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(booking.EAUTHNOTFOUND, "Could not find auth")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find auth: %w", err)
	}

	return a, nil
}

func updateAuth(
	ctx context.Context,
	tx *Tx,
	req booking.UpdateAuthRequest,
	attachEdges func(*Auth) (*Auth, error),
) (*Auth, error) {
	a, err := tx.Auth.
		UpdateOneID(req.ID).
		SetAccessToken(req.AccessToken).
		SetRefreshToken(req.RefreshToken).
		SetNillableExpiry(req.Expiry).
		Save(ctx)

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return nil, booking.Errorf(
			booking.EAUTHNOTFOUND,
			"Could not find auth with ID %d",
			req.ID,
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update auth: %w", err)
	}

	a, err = attachEdges(a)
	if err != nil {
		return nil, fmt.Errorf("failed to attach edges: %w", err)
	}

	return a, nil
}

// UpdateAuth updates an auth in the store.
func (s *authService) UpdateAuth(ctx context.Context, req booking.UpdateAuthRequest) booking.UpdateAuthResponse {
	panic("not implemented") // TODO: Implement
}

// DeleteAuth permanently deletes an authentication object from the system by ID. The
// parent user object is not removed.
func (s *authService) DeleteAuth(ctx context.Context, req booking.DeleteAuthRequest) booking.DeleteAuthResponse {
	panic("not implemented") // TODO: Implement
}

func (a *Auth) toModel() *booking.Auth {
	result := &booking.Auth{
		ID:           a.ID,
		UserID:       a.UserId,
		Source:       a.Source,
		SourceID:     a.SourceId,
		AccessToken:  *a.AccessToken,
		RefreshToken: *a.RefreshToken,
		Expiry:       a.Expiry,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
	if a.Edges.User != nil {
		result.User = a.Edges.User.toModel()
	}
	return result
}

func (a Auths) toModels() []*booking.Auth {
	var auths []*booking.Auth
	for _, v := range a {
		auths = append(auths, v.toModel())
	}
	return auths
}

func (u *User) toModel() *booking.User {
	result := &booking.User{
		ID:             u.ID,
		OrganizationID: u.OrganizationId,
		Name:           u.Name,
		Email:          u.Email,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
	if u.Edges.Auths != nil {
		result.Auths = Auths(u.Edges.Auths).toModels()
	}
	if u.Edges.Organization != nil {
		result.Organization = u.Edges.Organization.toModel()
	}
	return result
}
