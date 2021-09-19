package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func UnavailabilityMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.UnavailabilityServiceMiddleware {
	return func(next booking.UnavailabilityService) booking.UnavailabilityService {
		return unavailabilityMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type unavailabilityMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.UnavailabilityService
}

func (mw unavailabilityMetricsMiddleware) FindUnavailabilityByID(ctx context.Context, req booking.FindUnavailabilityByIDRequest) (res booking.FindUnavailabilityByIDResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_unavailability_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.UnavailabilityService.FindUnavailabilityByID(ctx, req)
	return
}

func (mw unavailabilityMetricsMiddleware) FindUnavailabilities(ctx context.Context, req booking.FindUnavailabilitiesRequest) (res booking.FindUnavailabilitiesResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_unavailabilities"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.UnavailabilityService.FindUnavailabilities(ctx, req)
	return
}

func (mw unavailabilityMetricsMiddleware) CreateUnavailability(ctx context.Context, req booking.CreateUnavailabilityRequest) (res booking.CreateUnavailabilityResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_unavailability"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.UnavailabilityService.CreateUnavailability(ctx, req)
	return
}

func (mw unavailabilityMetricsMiddleware) UpdateUnavailability(ctx context.Context, req booking.UpdateUnavailabilityRequest) (res booking.UpdateUnavailabilityResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_unavailability"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.UnavailabilityService.UpdateUnavailability(ctx, req)
	return
}

func (mw unavailabilityMetricsMiddleware) DeleteUnavailability(ctx context.Context, req booking.DeleteUnavailabilityRequest) (res booking.DeleteUnavailabilityResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_unavailability"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.UnavailabilityService.DeleteUnavailability(ctx, req)
	return
}
