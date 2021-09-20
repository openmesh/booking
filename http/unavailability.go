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

func (s *Server) registerUnavailabilityRoutes(r *mux.Router) {
	e := endpoint.MakeUnavailabilityEndpoints(s.UnavailabilityService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/resources/{resourceId}/unavailabilities/{id}").Handler(httptransport.NewServer(
		e.FindUnavailabilityByIDEndpoint,
		decodeFindUnavailabilityByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/resources/{resourceId}/unavailabilities").Handler(httptransport.NewServer(
		e.FindUnavailabilitiesEndpoint,
		decodeFindUnavailabilitiesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/resources/{resourceId}/unavailabilities").Handler(httptransport.NewServer(
		e.CreateUnavailabilityEndpoint,
		decodeCreateUnavailabilityRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/resources/{resourceId}/unavailabilities").Handler(httptransport.NewServer(
		e.UpdateUnavailabilityEndpoint,
		decodeUpdateUnavailabilityRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/resources/{resourceId}/unavailabilities/{id}").Handler(httptransport.NewServer(
		e.DeleteUnavailabilityEndpoint,
		decodeDeleteUnavailabilityRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindUnavailabilityByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindUnavailabilityByIDRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeFindUnavailabilitiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindUnavailabilitiesRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.CreateUnavailabilityRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.UpdateUnavailabilityRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.DeleteUnavailabilityRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}
