package http

import (
	"context"
	"net/http"

	"github.com/openmesh/booking"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/openmesh/booking/endpoint"
)

func (s *Server) registerReportRoutes(r *mux.Router) {
	e := endpoint.MakeReportEndpoints(s.ReportService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetRecentSalesReportEndpoint,
		decodeGetRecentSalesReportRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetUpcomingBookingsReportEndpoint,
		decodeGetUpcomingBookingsReportRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetBookingsActivityReportEndpoint,
		decodeGetBookingsActivityReportRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetTodaysBookingsReportEndpoint,
		decodeGetTodaysBookingsReportRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetTopResourcesReportEndpoint,
		decodeGetTopResourcesReportRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.GetTopEmployeesReportEndpoint,
		decodeGetTopEmployeesReportRequest,
		encodeResponse,
		options...,
	))
}

func decodeGetRecentSalesReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetRecentSalesReportRequest{}, nil
}

func decodeGetUpcomingBookingsReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetUpcomingBookingsReportRequest{}, nil
}

func decodeGetBookingsActivityReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetBookingsActivityReportRequest{}, nil
}

func decodeGetTodaysBookingsReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetTodaysBookingsReportRequest{}, nil
}

func decodeGetTopResourcesReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetTopResourcesReportRequest{}, nil
}

func decodeGetTopEmployeesReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.GetTopEmployeesReportRequest{}, nil
}
