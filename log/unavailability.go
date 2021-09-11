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

func (mw unavailabilityLoggingMiddleware) FindUnavailabilityByID(ctx context.Context, id int) (unavailability *booking.Unavailability, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_unavailability_by_id",
			"id", id,
			"unavailability", unavailability,
			"err", err,
		)
	}(time.Now())

	unavailability, err = mw.UnavailabilityService.FindUnavailabilityByID(ctx, id)
	return
}

func (mw unavailabilityLoggingMiddleware) FindUnavailabilities(ctx context.Context, filter booking.UnavailabilityFilter) (unavailabilities []*booking.Unavailability, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_unavailabilities",
			"filter", filter,
			"unavailabilities", unavailabilities,
			"totalItems", totalItems,
			"err", err,
		)
	}(time.Now())

	unavailabilities, totalItems, err = mw.UnavailabilityService.FindUnavailabilities(ctx, filter)
	return
}

func (mw unavailabilityLoggingMiddleware) CreateUnavailability(ctx context.Context, unavailability *booking.Unavailability) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_unavailability",
			"unavailability", unavailability,
			"err", err,
		)
	}(time.Now())

	err = mw.UnavailabilityService.CreateUnavailability(ctx, unavailability)
	return
}

func (mw unavailabilityLoggingMiddleware) UpdateUnavailability(ctx context.Context, id int, upd booking.UnavailabilityUpdate) (unavailability *booking.Unavailability, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_unavailability",
			"id", id,
			"update", upd,
			"unavailability", unavailability,
			"err", err,
		)
	}(time.Now())

	unavailability, err = mw.UnavailabilityService.UpdateUnavailability(ctx, id, upd)
	return
}

func (mw unavailabilityLoggingMiddleware) DeleteUnavailability(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_unavailability",
			"id", id,
			"err", err,
		)
	}(time.Now())

	err = mw.UnavailabilityService.DeleteUnavailability(ctx, id)
	return
}
