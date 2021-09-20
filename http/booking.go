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

func (s *Server) registerBookingRoutes(r *mux.Router) {
	e := endpoint.MakeBookingEndpoints(s.BookingService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/bookings/{id}").Handler(httptransport.NewServer(
		e.FindBookingByIDEndpoint,
		decodeFindBookingByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/bookings").Handler(httptransport.NewServer(
		e.FindBookingsEndpoint,
		decodeFindBookingsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/bookings").Handler(httptransport.NewServer(
		e.CreateBookingEndpoint,
		decodeCreateBookingRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/bookings/{id}").Handler(httptransport.NewServer(
		e.UpdateBookingEndpoint,
		decodeUpdateBookingRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/bookings/{id}").Handler(httptransport.NewServer(
		e.DeleteBookingEndpoint,
		decodeDeleteBookingRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindBookingByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindBookingByIDRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeFindBookingsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindBookingsRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.CreateBookingRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.UpdateBookingRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.DeleteBookingRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}
