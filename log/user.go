package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func UserLoggingMiddleware(logger log.Logger) booking.UserServiceMiddleware {
	return func(next booking.UserService) booking.UserService {
		return userLoggingMiddleware{logger, next}
	}
}

type userLoggingMiddleware struct {
	logger log.Logger
	booking.UserService
}

func (mw userLoggingMiddleware) FindUserByID(ctx context.Context, id int) (user *booking.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_user_by_id",
			"id", id,
			"user", user,
			"err", err,
		)
	}(time.Now())

	user, err = mw.UserService.FindUserByID(ctx, id)
	return
}

func (mw userLoggingMiddleware) FindUsers(ctx context.Context, filter booking.UserFilter) (users []*booking.User, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_users",
			"filter", filter,
			"users", users,
			"total_items", totalItems,
			"err", err,
		)
	}(time.Now())

	users, totalItems, err = mw.UserService.FindUsers(ctx, filter)
	return
}

func (mw userLoggingMiddleware) CreateUser(ctx context.Context, user *booking.User) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_user",
			"user", user,
			"err", err,
		)
	}(time.Now())

	err = mw.UserService.CreateUser(ctx, user)
	return
}

func (mw userLoggingMiddleware) UpdateUser(ctx context.Context, id int, upd booking.UserUpdate) (user *booking.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_user",
			"id", id,
			"update", upd,
			"user", user,
			"err", err,
		)
	}(time.Now())

	user, err = mw.UserService.UpdateUser(ctx, id, upd)
	return
}

func (mw userLoggingMiddleware) DeleteUser(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_user",
			"id", id,
			"err", err,
		)
	}(time.Now())

	err = mw.UserService.DeleteUser(ctx, id)
	return
}
