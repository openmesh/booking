package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func BookingLoggingMiddleware(logger log.Logger) booking.BookingServiceMiddleware {
	return func(next booking.BookingService) booking.BookingService {
		return bookingLoggingMiddleware{logger, next}
	}
}

type bookingLoggingMiddleware struct {
	logger log.Logger
	booking.BookingService
}

func (mw bookingLoggingMiddleware) FindBookingByID(ctx context.Context, id int) (booking *booking.Booking, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_booking_by_id",
			"id", id,
			"booking", booking,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	booking, err = mw.BookingService.FindBookingByID(ctx, id)
	return
}

func (mw bookingLoggingMiddleware) FindBookings(ctx context.Context, filter booking.BookingFilter) (bookings []*booking.Booking, totalItems int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_bookings",
			"filter", filter,
			"bookings", bookings,
			"total_items", totalItems,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	bookings, totalItems, err = mw.BookingService.FindBookings(ctx, filter)
	return
}

func (mw bookingLoggingMiddleware) CreateBooking(ctx context.Context, booking *booking.Booking) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_booking",
			"booking", booking,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.BookingService.CreateBooking(ctx, booking)
	return
}

func (mw bookingLoggingMiddleware) UpdateBooking(ctx context.Context, id int, upd booking.BookingUpdate) (booking *booking.Booking, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_booking",
			"id", id,
			"update", upd,
			"booking", booking,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	booking, err = mw.BookingService.UpdateBooking(ctx, id, upd)
	return
}

func (mw bookingLoggingMiddleware) DeleteBooking(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_booking",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.BookingService.DeleteBooking(ctx, id)
	return
}
