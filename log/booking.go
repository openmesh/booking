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

func (mw bookingLoggingMiddleware) FindBookingByID(ctx context.Context, req booking.FindBookingByIDRequest) (res booking.FindBookingByIDResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_booking_by_id",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.BookingService.FindBookingByID(ctx, req)
	return
}

func (mw bookingLoggingMiddleware) FindBookings(ctx context.Context, req booking.FindBookingsRequest) (res booking.FindBookingsResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_bookings",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.BookingService.FindBookings(ctx, req)
	return
}

func (mw bookingLoggingMiddleware) CreateBooking(ctx context.Context, req booking.CreateBookingRequest) (res booking.CreateBookingResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_booking",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.BookingService.CreateBooking(ctx, req)
	return
}

func (mw bookingLoggingMiddleware) UpdateBooking(ctx context.Context, req booking.UpdateBookingRequest) (res booking.UpdateBookingResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_booking",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.BookingService.UpdateBooking(ctx, req)
	return
}

func (mw bookingLoggingMiddleware) DeleteBooking(ctx context.Context, req booking.DeleteBookingRequest) (res booking.DeleteBookingResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_booking",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.BookingService.DeleteBooking(ctx, req)
	return
}
