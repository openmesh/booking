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

func (mw resourceLoggingMiddleware) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) (res *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resource_by_id",
			"request", req,
			"resource", res,
			"err", err,
		)
	}(time.Now())

	res, err = mw.ResourceService.FindResourceByID(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) FindResources(ctx context.Context, req booking.FindResourcesRequest) (resources []*booking.Resource, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resources",
			"request", req,
			"resources", resources,
			"total_items", totalItems,
			"err", err,
		)
	}(time.Now())

	resources, totalItems, err = mw.ResourceService.FindResources(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) CreateResource(ctx context.Context, req booking.CreateResourceRequest) (result *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_resource",
			"request", req,
			"err", err,
		)
	}(time.Now())

	result, err = mw.ResourceService.CreateResource(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) (resource *booking.Resource, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_resource",
			"request", req,
			"resource", resource,
			"err", err,
		)
	}(time.Now())

	resource, err = mw.ResourceService.UpdateResource(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_resource",
			"request", req,
			"err", err,
		)
	}(time.Now())

	err = mw.ResourceService.DeleteResource(ctx, req)
	return
}
