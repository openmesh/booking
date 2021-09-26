package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func OAuthMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.OAuthServiceMiddleware {
	return func(next booking.OAuthService) booking.OAuthService {
		return oauthMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type oauthMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.OAuthService
}

func (mw oauthMetricsMiddleware) GetRedirectURL(ctx context.Context, req booking.GetRedirectURLRequest) (res booking.GetRedirectURLResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_redirect_url"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.OAuthService.GetRedirectURL(ctx, req)
	return
}

func (mw oauthMetricsMiddleware) HandleCallback(ctx context.Context, req booking.HandleCallbackRequest) (res booking.HandleCallbackResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "handle_callback"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.OAuthService.HandleCallback(ctx, req)
	return
}
