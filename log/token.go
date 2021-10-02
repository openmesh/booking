package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func TokenLoggingMiddleware(logger log.Logger) booking.TokenServiceMiddleware {
	return func(next booking.TokenService) booking.TokenService {
		return tokenLoggingMiddleware{logger, next}
	}
}

type tokenLoggingMiddleware struct {
	logger log.Logger
	booking.TokenService
}

func (mw tokenLoggingMiddleware) CreateToken(ctx context.Context, req booking.CreateTokenRequest) (res booking.CreateTokenResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_token",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.TokenService.CreateToken(ctx, req)
	return
}

func (mw tokenLoggingMiddleware) FindTokens(ctx context.Context, req booking.FindTokensRequest) (res booking.FindTokensResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_tokens",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.TokenService.FindTokens(ctx, req)
	return
}
