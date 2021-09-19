package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func UnavailabilityLoggingMiddleware(logger log.Logger) booking.UnavailabilityServiceMiddleware {
	return func(next booking.UnavailabilityService) booking.UnavailabilityService {
		return unavailabilityLoggingMiddleware{logger, next}
	}
}

type unavailabilityLoggingMiddleware struct {
	logger log.Logger
	booking.UnavailabilityService
}

func (mw unavailabilityLoggingMiddleware) FindUnavailabilityByID(ctx context.Context, req booking.FindUnavailabilityByIDRequest) (res booking.FindUnavailabilityByIDResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_unavailability_by_id",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.UnavailabilityService.FindUnavailabilityByID(ctx, req)
	return
}

func (mw unavailabilityLoggingMiddleware) FindUnavailabilities(ctx context.Context, req booking.FindUnavailabilitiesRequest) (res booking.FindUnavailabilitiesResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_unavailabilities",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.UnavailabilityService.FindUnavailabilities(ctx, req)
	return
}

func (mw unavailabilityLoggingMiddleware) CreateUnavailability(ctx context.Context, req booking.CreateUnavailabilityRequest) (res booking.CreateUnavailabilityResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_unavailability",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.UnavailabilityService.CreateUnavailability(ctx, req)
	return
}

func (mw unavailabilityLoggingMiddleware) UpdateUnavailability(ctx context.Context, req booking.UpdateUnavailabilityRequest) (res booking.UpdateUnavailabilityResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_unavailability",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.UnavailabilityService.UpdateUnavailability(ctx, req)
	return
}

func (mw unavailabilityLoggingMiddleware) DeleteUnavailability(ctx context.Context, req booking.DeleteUnavailabilityRequest) (res booking.DeleteUnavailabilityResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_unavailability",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.UnavailabilityService.DeleteUnavailability(ctx, req)
	return
}
