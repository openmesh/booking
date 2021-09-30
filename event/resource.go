package event

import (
	"context"

	"github.com/openmesh/booking"
)

func ResourceEventMiddleware(eventService booking.EventService) booking.ResourceServiceMiddleware {
	return func(next booking.ResourceService) booking.ResourceService {
		return resourceEventMiddleware{eventService, next}
	}
}

type resourceEventMiddleware struct {
	booking.EventService
	booking.ResourceService
}

// FindResourceByID retrieves a single resource by ID along with associated availabilities.
// Returns ENOTFOUND if resource does not exist or user does not have
// permission to view it.
func (mw resourceEventMiddleware) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) booking.FindResourceByIDResponse {
	return mw.ResourceService.FindResourceByID(ctx, req)
}

// FindResources retrieves a lit of resources based on a filter. Only returns
// resources that accessible to the user. Also returns a count of total matching bookings
// which may be different from the number of returned bookings if the "Limit" field is set.
func (mw resourceEventMiddleware) FindResources(ctx context.Context, req booking.FindResourcesRequest) booking.FindResourcesResponse {
	return mw.ResourceService.FindResources(ctx, req)
}

// CreateResource creates a new resource and assigns the current user as the owner.
// Returns the created Resource.
func (mw resourceEventMiddleware) CreateResource(ctx context.Context, req booking.CreateResourceRequest) (res booking.CreateResourceResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeResourceCreated,
			Payload: booking.ResourceCreatedPayload{Resource: res.Resource},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.ResourceService.CreateResource(ctx, req)
	return
}

// UpdateResource updates an existing resource by ID. Only the resource owner can update a
// resource. Returns the new resource state even if there was an error during update.
//
// Returns ENOTFOUND if the resource does not exist or the user does not have
// permission to update it.
func (mw resourceEventMiddleware) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) (res booking.UpdateResourceResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeResourceUpdated,
			Payload: booking.ResourceUpdatedPayload{Resource: res.Resource},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.ResourceService.UpdateResource(ctx, req)
	return
}

// DeleteResource permanently removes a resource by ID. Only the resource owner may delete a
// resource. Returns ENOTFOUND if the resource does not exist or the user does not have
// permission to delete it.
func (mw resourceEventMiddleware) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) (res booking.DeleteResourceResponse) {
	defer func() {
		if res.Err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeResourceDeleted,
			Payload: booking.ResourceDeletedPayload{ID: req.ID},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	res = mw.ResourceService.DeleteResource(ctx, req)
	return
}
