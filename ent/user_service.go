package ent

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent/user"
)

type userService struct {
	client *Client
}

func NewUserService(client *Client) *userService {
	return &userService{
		client,
	}
}

// Retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *userService) FindUserByID(ctx context.Context, id int) (*booking.User, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	u, err := tx.User.
		Query().
		WithOrganization().
		Where(user.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return u.toModel(), nil
}

// Retrieves a list of users by filter. Also returns total count of matching
// users which may differ from returned results if filter.Limit is specified.
func (s *userService) FindUsers(ctx context.Context, filter booking.UserFilter) ([]*booking.User, int, error) {
	panic("not implemented") // TODO: Implement
}

// Creates a new user. This is only used for testing since users are typically
// created during the OAuth creation process in AuthService.CreateAuth().
func (s *userService) CreateUser(ctx context.Context, user *booking.User) error {
	panic("not implemented") // TODO: Implement
}

// Updates a user object. Returns EUNAUTHORIZED if current user is not
// the user that is being updated. Returns ENOTFOUND if user does not exist.
func (s *userService) UpdateUser(ctx context.Context, id int, upd booking.UserUpdate) (*booking.User, error) {
	panic("not implemented") // TODO: Implement
}

// Permanently deletes a user and all owned dials. Returns EUNAUTHORIZED
// if current user is not the user being deleted. Returns ENOTFOUND if
// user does not exist.
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
