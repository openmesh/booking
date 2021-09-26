package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func AuthLoggingMiddleware(logger log.Logger) booking.AuthServiceMiddleware {
	return func(next booking.AuthService) booking.AuthService {
		return authLoggingMiddleware{logger, next}
	}
}

type authLoggingMiddleware struct {
	logger log.Logger
	booking.AuthService
}

func (mw authLoggingMiddleware) FindAuthByID(ctx context.Context, req booking.FindAuthByIDRequest) (res booking.FindAuthByIDResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_auth_by_id",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.AuthService.FindAuthByID(ctx, req)
	return
}

func (mw authLoggingMiddleware) FindAuths(ctx context.Context, req booking.FindAuthsRequest) (res booking.FindAuthsResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_auths",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.AuthService.FindAuths(ctx, req)
	return
}

func (mw authLoggingMiddleware) CreateAuth(ctx context.Context, req booking.CreateAuthRequest) (res booking.CreateAuthResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_auth",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.AuthService.CreateAuth(ctx, req)
	return
}

func (mw authLoggingMiddleware) UpdateAuth(ctx context.Context, req booking.UpdateAuthRequest) (res booking.UpdateAuthResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_auth",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.AuthService.UpdateAuth(ctx, req)
	return
}

func (mw authLoggingMiddleware) DeleteAuth(ctx context.Context, req booking.DeleteAuthRequest) (res booking.DeleteAuthResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_auth",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.AuthService.DeleteAuth(ctx, req)
	return
}
