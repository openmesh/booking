package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func OAuthLoggingMiddleware(logger log.Logger) booking.OAuthServiceMiddleware {
	return func(next booking.OAuthService) booking.OAuthService {
		return oauthLoggingMiddleware{logger, next}
	}
}

type oauthLoggingMiddleware struct {
	logger log.Logger
	booking.OAuthService
}

func (mw oauthLoggingMiddleware) GetRedirectURL(ctx context.Context, req booking.GetRedirectURLRequest) (res booking.GetRedirectURLResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_redirect_url",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.OAuthService.GetRedirectURL(ctx, req)
	return
}

func (mw oauthLoggingMiddleware) HandleCallback(ctx context.Context, req booking.HandleCallbackRequest) (res booking.HandleCallbackResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "handle_callback",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.OAuthService.HandleCallback(ctx, req)
	return
}
