package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func ReportLoggingMiddleware(logger log.Logger) booking.ReportServiceMiddleware {
	return func(next booking.ReportService) booking.ReportService {
		return reportLoggingMiddleware{logger, next}
	}
}

type reportLoggingMiddleware struct {
	logger log.Logger
	booking.ReportService
}

func (mw reportLoggingMiddleware) GetRecentSalesReport(ctx context.Context, req booking.GetRecentSalesReportRequest) (res booking.GetRecentSalesReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_recent_sales_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetRecentSalesReport(ctx, req)
	return
}

func (mw reportLoggingMiddleware) GetUpcomingBookingsReport(ctx context.Context, req booking.GetUpcomingBookingsReportRequest) (res booking.GetUpcomingBookingsReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_upcoming_bookings_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetUpcomingBookingsReport(ctx, req)
	return
}

func (mw reportLoggingMiddleware) GetBookingsActivityReport(ctx context.Context, req booking.GetBookingsActivityReportRequest) (res booking.GetBookingsActivityReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_bookings_activity_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetBookingsActivityReport(ctx, req)
	return
}

func (mw reportLoggingMiddleware) GetTodaysBookingsReport(ctx context.Context, req booking.GetTodaysBookingsReportRequest) (res booking.GetTodaysBookingsReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_todays_bookings_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetTodaysBookingsReport(ctx, req)
	return
}

func (mw reportLoggingMiddleware) GetTopResourcesReport(ctx context.Context, req booking.GetTopResourcesReportRequest) (res booking.GetTopResourcesReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_top_resources_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetTopResourcesReport(ctx, req)
	return
}

func (mw reportLoggingMiddleware) GetTopEmployeesReport(ctx context.Context, req booking.GetTopEmployeesReportRequest) (res booking.GetTopEmployeesReportResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_top_employees_report",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.ReportService.GetTopEmployeesReport(ctx, req)
	return
}
