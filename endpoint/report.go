package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/booking"
)

// ReportEndpoints collects all the endpoints that compose a booking.ReportService.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type ReportEndpoints struct {
	GetRecentSalesReportEndpoint      endpoint.Endpoint
	GetUpcomingBookingsReportEndpoint endpoint.Endpoint
	GetBookingsActivityReportEndpoint endpoint.Endpoint
	GetTodaysBookingsReportEndpoint   endpoint.Endpoint
	GetTopResourcesReportEndpoint     endpoint.Endpoint
	GetTopEmployeesReportEndpoint     endpoint.Endpoint
}

// MakeReportEndpoints returns a ReportEndpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeReportEndpoints(s booking.ReportService) ReportEndpoints {
	return ReportEndpoints{
		GetRecentSalesReportEndpoint:      MakeGetRecentSalesReportEndpoint(s),
		GetUpcomingBookingsReportEndpoint: MakeGetUpcomingBookingsReportEndpoint(s),
		GetBookingsActivityReportEndpoint: MakeGetBookingsActivityReportEndpoint(s),
		GetTodaysBookingsReportEndpoint:   MakeGetTodaysBookingsReportEndpoint(s),
		GetTopResourcesReportEndpoint:     MakeGetTopResourcesReportEndpoint(s),
		GetTopEmployeesReportEndpoint:     MakeGetTopEmployeesReportEndpoint(s),
	}
}

// MakeGetRecentSalesReportEndpoint returns an endpoint via the passed service.
func MakeGetRecentSalesReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetRecentSalesReport(ctx, r.(booking.GetRecentSalesReportRequest)), nil
	}
}

// MakeGetUpcomingBookingsReportEndpoint returns an endpoint via the passed service.
func MakeGetUpcomingBookingsReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetUpcomingBookingsReport(ctx, r.(booking.GetUpcomingBookingsReportRequest)), nil
	}
}

// MakeGetBookingsActivityReportEndpoint returns an endpoint via the passed service.
func MakeGetBookingsActivityReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetBookingsActivityReport(ctx, r.(booking.GetBookingsActivityReportRequest)), nil
	}
}

// MakeGetTodaysBookingsReportEndpoint returns an endpoint via the passed service.
func MakeGetTodaysBookingsReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetTodaysBookingsReport(ctx, r.(booking.GetTodaysBookingsReportRequest)), nil
	}
}

// MakeGetTopResourcesReportEndpoint returns an endpoint via the passed service.
func MakeGetTopResourcesReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetTopResourcesReport(ctx, r.(booking.GetTopResourcesReportRequest)), nil
	}
}

// MakeGetTopEmployeesReportEndpoint returns an endpoint via the passed service.
func MakeGetTopEmployeesReportEndpoint(s booking.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.GetTopEmployeesReport(ctx, r.(booking.GetTopEmployeesReportRequest)), nil
	}
}
