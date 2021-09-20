package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func BookingMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.BookingServiceMiddleware {
	return func(next booking.BookingService) booking.BookingService {
		return bookingMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type bookingMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.BookingService
}


func (mw bookingMetricsMiddleware) FindBookingByID(ctx context.Context, req booking.FindBookingByIDRequest) (res booking.FindBookingByIDResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_booking_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.BookingService.FindBookingByID(ctx, req)
	return
}

func (mw bookingMetricsMiddleware) FindBookings(ctx context.Context, req booking.FindBookingsRequest) (res booking.FindBookingsResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_bookings"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.BookingService.FindBookings(ctx, req)
	return
}

func (mw bookingMetricsMiddleware) CreateBooking(ctx context.Context, req booking.CreateBookingRequest) (res booking.CreateBookingResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_booking"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.BookingService.CreateBooking(ctx, req)
	return
}

func (mw bookingMetricsMiddleware) UpdateBooking(ctx context.Context, req booking.UpdateBookingRequest) (res booking.UpdateBookingResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_booking"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.BookingService.UpdateBooking(ctx, req)
	return
}

func (mw bookingMetricsMiddleware) DeleteBooking(ctx context.Context, req booking.DeleteBookingRequest) (res booking.DeleteBookingResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_booking"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.BookingService.DeleteBooking(ctx, req)
	return
}
