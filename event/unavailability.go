package event

import (
	"context"

	"github.com/openmesh/booking"
)

func UnavailabilityEventMiddleware(eventService booking.EventService) booking.UnavailabilityServiceMiddleware {
	return func(next booking.UnavailabilityService) booking.UnavailabilityService {
		return unavailabilityEventMiddleware{eventService, next}
	}
}

type unavailabilityEventMiddleware struct {
	booking.EventService
	booking.UnavailabilityService
}

// Retrieves a single unavailability by ID along with the associated resource.
// Returns ENOTFOUND if unavailability does not exist or user does not have
// permission to view it.
func (mw unavailabilityEventMiddleware) FindUnavailabilityByID(ctx context.Context, req booking.FindUnavailabilityByIDRequest) booking.FindUnavailabilityByIDResponse {
	return mw.UnavailabilityService.FindUnavailabilityByID(ctx, req)
}

// Retreives a list of unavailabilities based on a filter. Only returns
// unavailabilities that are accessible to the user. Also returns a count of
// total matching unavailabilities which may be different from the number of
// returned unavailabilities if the "Limit" field is set.
func (mw unavailabilityEventMiddleware) FindUnavailabilities(ctx context.Context, req booking.FindUnavailabilitiesRequest) booking.FindUnavailabilitiesResponse {
	return mw.UnavailabilityService.FindUnavailabilities(ctx, req)
}

// Create a new unavailability and assigns the current user as the owner.
func (mw unavailabilityEventMiddleware) CreateUnavailability(ctx context.Context, req booking.CreateUnavailabilityRequest) (res booking.CreateUnavailabilityResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeUnavailabilityCreated,
			Payload: booking.UnavailabilityCreatedPayload{Unavailability: res.Unavailability},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.UnavailabilityService.CreateUnavailability(ctx, req)
	return
}

// Updates an existing unavailbility by ID. Only the unavailability owner can
// update an unavailability. Returns the new unavailability state even if
// there was an error during update.
//
// Returns ENOTFOUND if the unavailability does not exist or the user does not
// have permission to update it.
func (mw unavailabilityEventMiddleware) UpdateUnavailability(ctx context.Context, req booking.UpdateUnavailabilityRequest) (res booking.UpdateUnavailabilityResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeUnavailabilityUpdated,
			Payload: booking.UnavailabilityUpdatedPayload{Unavailability: res.Unavailability},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.UnavailabilityService.UpdateUnavailability(ctx, req)
	return
}

// Permanently removes a unavailability by ID. Only the unavailability owner
// may delete a unavailability. Returns ENOTFOUND if the unavailability does
// not exist or the user does not have permission to delete it.
func (mw unavailabilityEventMiddleware) DeleteUnavailability(ctx context.Context, req booking.DeleteUnavailabilityRequest) (res booking.DeleteUnavailabilityResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeUnavailabilityDeleted,
			Payload: booking.UnavailabilityDeletedPayload{ID: req.ID},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.UnavailabilityService.DeleteUnavailability(ctx, req)
	return
}
