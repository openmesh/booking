package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func TokenMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.TokenServiceMiddleware {
	return func(next booking.TokenService) booking.TokenService {
		return tokenMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type tokenMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.TokenService
}

func (mw tokenMetricsMiddleware) CreateToken(ctx context.Context, req booking.CreateTokenRequest) (res booking.CreateTokenResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_token"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.TokenService.CreateToken(ctx, req)
	return
}

func (mw tokenMetricsMiddleware) FindTokens(ctx context.Context, req booking.FindTokensRequest) (res booking.FindTokensResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_tokens"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.TokenService.FindTokens(ctx, req)
	return
}
