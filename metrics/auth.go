package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func AuthMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.AuthServiceMiddleware {
	return func(next booking.AuthService) booking.AuthService {
		return authMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type authMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.AuthService
}

func (mw authMetricsMiddleware) FindAuthByID(ctx context.Context, req booking.FindAuthByIDRequest) (res booking.FindAuthByIDResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_auth_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.AuthService.FindAuthByID(ctx, req)
	return
}

func (mw authMetricsMiddleware) FindAuths(ctx context.Context, req booking.FindAuthsRequest) (res booking.FindAuthsResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_auths"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.AuthService.FindAuths(ctx, req)
	return
}

func (mw authMetricsMiddleware) CreateAuth(ctx context.Context, req booking.CreateAuthRequest) (res booking.CreateAuthResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_auth"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.AuthService.CreateAuth(ctx, req)
	return
}

func (mw authMetricsMiddleware) UpdateAuth(ctx context.Context, req booking.UpdateAuthRequest) (res booking.UpdateAuthResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_auth"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.AuthService.UpdateAuth(ctx, req)
	return
}

func (mw authMetricsMiddleware) DeleteAuth(ctx context.Context, req booking.DeleteAuthRequest) (res booking.DeleteAuthResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_auth"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.AuthService.DeleteAuth(ctx, req)
	return
}
