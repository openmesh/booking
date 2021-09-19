package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func ResourceMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.ResourceServiceMiddleware {
	return func(next booking.ResourceService) booking.ResourceService {
		return resourceMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type resourceMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.ResourceService
}

func (mw resourceMetricsMiddleware) FindResourceByID(ctx context.Context, req booking.FindResourceByIDRequest) (res booking.FindResourceByIDResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_resource_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	res = mw.ResourceService.FindResourceByID(ctx, req)
	return
}
