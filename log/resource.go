package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func ResourceLoggingMiddleware(logger log.Logger) booking.ResourceServiceMiddleware {
	return func(next booking.ResourceService) booking.ResourceService {
		return resourceLoggingMiddleware{logger, next}
	}
}

type resourceLoggingMiddleware struct {
	logger log.Logger
	booking.ResourceService
}

func (mw resourceLoggingMiddleware) FindResourceByID(ctx context.Context, id int) (resource *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resource_by_id",
			"id", id,
			"resource", resource,
			"err", err,
		)
	}(time.Now())

	resource, err = mw.ResourceService.FindResourceByID(ctx, id)
	return
}

func (mw resourceLoggingMiddleware) FindResources(ctx context.Context, filter booking.ResourceFilter) (resources []*booking.Resource, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resources",
			"filter", filter,
			"resources", resources,
			"total_items", totalItems,
			"err", err,
		)
	}(time.Now())

	resources, totalItems, err = mw.ResourceService.FindResources(ctx, filter)
	return
}

func (mw resourceLoggingMiddleware) CreateResource(ctx context.Context, resource *booking.Resource) (result *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_resource",
			"resource", resource,
			"err", err,
		)
	}(time.Now())

	resource, err = mw.ResourceService.CreateResource(ctx, resource)
	return
}

func (mw resourceLoggingMiddleware) UpdateResource(ctx context.Context, id int, upd booking.ResourceUpdate) (resource *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_resource",
			"id", id,
			"update", upd,
			"resource", resource,
			"err", err,
		)
	}(time.Now())

	resource, err = mw.ResourceService.UpdateResource(ctx, id, upd)
	return
}

func (mw resourceLoggingMiddleware) DeleteResource(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_resource",
			"id", id,
			"err", err,
		)
	}(time.Now())

	err = mw.ResourceService.DeleteResource(ctx, id)
	return
}
