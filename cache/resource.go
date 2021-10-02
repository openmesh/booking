package cache

import (
	"context"
	"fmt"

	"github.com/openmesh/booking"
)

func ResourceCacheMiddleware(c booking.Cache) booking.ResourceServiceMiddleware {
	return func(next booking.ResourceService) booking.ResourceService {
		return resourceCacheMiddleware{c, next}
	}
}

type resourceCacheMiddleware struct {
	booking.Cache
	booking.ResourceService
}

// FindResourceByID retrieves a single resource by ID along with associated availabilities.
// Returns ENOTFOUND if resource does not exist or user does not have
// permission to view it.
func (mw resourceCacheMiddleware) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) booking.FindResourceByIDResponse {
	key := booking.CacheKey(ctx, "find_resource_by_id", req)
	var res booking.FindResourceByIDResponse
	err := mw.Cache.Get(ctx, key, &res)
	if err != nil {
		res = mw.ResourceService.FindResourceByID(ctx, req)
		if res.Err != nil {
			return res
		}
		mw.Cache.Set(ctx, key, res)
		return res
	}
	return res
}

// FindResources retrieves a lit of resources based on a filter. Only returns
// resources that accessible to the user. Also returns a count of total matching bookings
// which may be different from the number of returned bookings if the "Limit" field is set.
func (mw resourceCacheMiddleware) FindResources(ctx context.Context, req booking.FindResourcesRequest) booking.FindResourcesResponse {
	key := booking.CacheKey(ctx, "find_resources", req)
	var res booking.FindResourcesResponse
	err := mw.Cache.Get(ctx, key, &res)
	if err != nil {
		res = mw.ResourceService.FindResources(ctx, req)
		if res.Err != nil {
			return res
		}
		mw.Cache.Set(ctx, key, res)
		return res
	}
	return res
}

// CreateResource creates a new resource and assigns the current user as the owner.
// Returns the created Resource.
func (mw resourceCacheMiddleware) CreateResource(ctx context.Context, req booking.CreateResourceRequest) booking.CreateResourceResponse {
	match := booking.CacheKey(ctx, "find_resources", "*")
	err := mw.Cache.RemoveMany(ctx, match)
	if err != nil {
		return booking.CreateResourceResponse{
			Resource: nil,
			Err:      fmt.Errorf("could not scan cache keys: %w", err),
		}
	}

	return mw.ResourceService.CreateResource(ctx, req)
}

// UpdateResource updates an existing resource by ID. Only the resource owner can update a
// resource. Returns the new resource state even if there was an error during update.
//
// Returns ENOTFOUND if the resource does not exist or the user does not have
// permission to update it.
func (mw resourceCacheMiddleware) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) booking.UpdateResourceResponse {
	match := booking.CacheKey(ctx, "find_resources", "*")
	err := mw.Cache.RemoveMany(ctx, match)
	if err != nil {
		return booking.UpdateResourceResponse{
			Resource: nil,
			Err:      fmt.Errorf("could not remove cache keys: %w", err),
		}
	}

	key := booking.CacheKey(ctx, "find_resource_by_id", booking.FindResourceByIDRequest{ID: req.ID})
	err = mw.Cache.Remove(ctx, key)
	if err != nil {
		return booking.UpdateResourceResponse{
			Resource: nil,
			Err:      fmt.Errorf("could not remove cache key: %w", err),
		}
	}

	return mw.ResourceService.UpdateResource(ctx, req)
}

// DeleteResource permanently removes a resource by ID. Only the resource owner may delete a
// resource. Returns ENOTFOUND if the resource does not exist or the user does not have
// permission to delete it.
func (mw resourceCacheMiddleware) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) booking.DeleteResourceResponse {
	panic("not implemented") // TODO: Implement
}
