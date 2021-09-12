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

func (mw authLoggingMiddleware) FindAuthByID(ctx context.Context, id int) (auth *booking.Auth, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_auth_by_id",
			"id", id,
			"auth", auth,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	auth, err = mw.AuthService.FindAuthByID(ctx, id)
	return
}

func (mw authLoggingMiddleware) FindAuths(ctx context.Context, filter booking.AuthFilter) (auths []*booking.Auth, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_auths",
			"filter", filter,
			"auths", auths,
			"total_items", totalItems,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	auths, totalItems, err = mw.AuthService.FindAuths(ctx, filter)
	return
}

func (mw authLoggingMiddleware) CreateAuth(ctx context.Context, auth *booking.Auth) (result *booking.Auth, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_auth",
			"auth", auth,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	result, err = mw.AuthService.CreateAuth(ctx, auth)
	return
}

func (mw authLoggingMiddleware) DeleteAuth(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_auth",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return
}
