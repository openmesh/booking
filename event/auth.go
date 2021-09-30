package event

import (
	"context"

	"github.com/openmesh/booking"
)

func AuthEventMiddleware(eventService booking.EventService) booking.AuthServiceMiddleware {
	return func(next booking.AuthService) booking.AuthService {
		return authEventMiddleware{eventService, next}
	}
}

type authEventMiddleware struct {
	booking.EventService
	booking.AuthService
}

// FindAuthByID looks up an authentication object by ID along with the associated user.
// Returns ENOTFOUND if ID does not exist.
func (mw authEventMiddleware) FindAuthByID(ctx context.Context, req booking.FindAuthByIDRequest) booking.FindAuthByIDResponse {
	return mw.AuthService.FindAuthByID(ctx, req)
}

// FindAuths retrieves authentication objects based on a filter. Also returns the total
// number of objects that match the filter. This may differ from the returned
// object count if the Limit field is set.
func (mw authEventMiddleware) FindAuths(ctx context.Context, req booking.FindAuthsRequest) booking.FindAuthsResponse {
	return mw.AuthService.FindAuths(ctx, req)
}

// CreateAuth creates a new authentication object If a User is attached to auth, then the
// auth object is linked to an existing user. Otherwise a new user object is
// created.
//
// Returns the created Auth.
func (mw authEventMiddleware) CreateAuth(ctx context.Context, req booking.CreateAuthRequest) (res booking.CreateAuthResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeAuthCreated,
			Payload: booking.AuthCreatedPayload{Auth: res.Auth},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.AuthService.CreateAuth(ctx, req)
	return
}

// UpdateAuth updates the refresh token value for an auth.
func (mw authEventMiddleware) UpdateAuth(ctx context.Context, req booking.UpdateAuthRequest) (res booking.UpdateAuthResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeAuthUpdated,
			Payload: booking.AuthCreatedPayload{Auth: res.Auth},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.AuthService.UpdateAuth(ctx, req)
	return
}

// DeleteAuth permanently deletes an authentication object from the system by ID. The
// parent user object is not removed.
func (mw authEventMiddleware) DeleteAuth(ctx context.Context, req booking.DeleteAuthRequest) (res booking.DeleteAuthResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeAuthDeleted,
			Payload: booking.AuthDeletedPayload{ID: req.ID},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.AuthService.DeleteAuth(ctx, req)
	return
}
