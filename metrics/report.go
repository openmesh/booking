package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/booking"
)

func ReportMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) booking.ReportServiceMiddleware {
	return func(next booking.ReportService) booking.ReportService {
		return reportMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type reportMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	booking.ReportService
}

func (mw reportMetricsMiddleware) GetRecentSalesReport(ctx context.Context, req booking.GetRecentSalesReportRequest) (res booking.GetRecentSalesReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_recent_sales_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetRecentSalesReport(ctx, req)
	return
}

func (mw reportMetricsMiddleware) GetUpcomingBookingsReport(ctx context.Context, req booking.GetUpcomingBookingsReportRequest) (res booking.GetUpcomingBookingsReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_upcoming_bookings_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetUpcomingBookingsReport(ctx, req)
	return
}

func (mw reportMetricsMiddleware) GetBookingsActivityReport(ctx context.Context, req booking.GetBookingsActivityReportRequest) (res booking.GetBookingsActivityReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_bookings_activity_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetBookingsActivityReport(ctx, req)
	return
}

func (mw reportMetricsMiddleware) GetTodaysBookingsReport(ctx context.Context, req booking.GetTodaysBookingsReportRequest) (res booking.GetTodaysBookingsReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_todays_bookings_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetTodaysBookingsReport(ctx, req)
	return
}

func (mw reportMetricsMiddleware) GetTopResourcesReport(ctx context.Context, req booking.GetTopResourcesReportRequest) (res booking.GetTopResourcesReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_top_resources_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetTopResourcesReport(ctx, req)
	return
}

func (mw reportMetricsMiddleware) GetTopEmployeesReport(ctx context.Context, req booking.GetTopEmployeesReportRequest) (res booking.GetTopEmployeesReportResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_top_employees_report"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.ReportService.GetTopEmployeesReport(ctx, req)
	return
}
