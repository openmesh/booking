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

func (mw resourceLoggingMiddleware) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) (res booking.FindResourceByIDResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resource_by_id",
			"request", req,
			"resource", res.Resource,
			"err", res.Err,
		)
	}(time.Now())
	res = mw.ResourceService.FindResourceByID(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) FindResources(ctx context.Context, req booking.FindResourcesRequest) (res booking.FindResourcesResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_resources",
			"request", req,
			"resources", res.Resources,
			"total_items", res.TotalItems,
			"err", res.Err,
		)
	}(time.Now())
	res = mw.ResourceService.FindResources(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) CreateResource(ctx context.Context, req booking.CreateResourceRequest) (res booking.CreateResourceResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_resource",
			"request", req,
			"resource", res.Resource,
			"err", res.Err,
		)
	}(time.Now())
	res = mw.ResourceService.CreateResource(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) UpdateResource(ctx context.Context, req booking.UpdateResourceRequest) (res booking.UpdateResourceResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_resource",
			"request", req,
			"resource", res.Resource,
			"err", res.Err,
		)
	}(time.Now())
	res = mw.ResourceService.UpdateResource(ctx, req)
	return
}

func (mw resourceLoggingMiddleware) DeleteResource(ctx context.Context, req booking.DeleteResourceRequest) (res booking.DeleteResourceResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_resource",
			"request", req,
			"err", res.Err,
		)
	}(time.Now())
	res = mw.ResourceService.DeleteResource(ctx, req)
	return
}
