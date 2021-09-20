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

func (s *Server) registerResourceRoutes(r *mux.Router) {
	e := endpoint.MakeResourceEndpoints(s.ResourceService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/resources/{id}").Handler(httptransport.NewServer(
		e.FindResourceByIDEndpoint,
		decodeFindResourceByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/resources").Handler(httptransport.NewServer(
		e.FindResourcesEndpoint,
		decodeFindResourcesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/resources").Handler(httptransport.NewServer(
		e.CreateResourceEndpoint,
		decodeCreateResourceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/resources/{id}").Handler(httptransport.NewServer(
		e.UpdateResourceEndpoint,
		decodeUpdateResourceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/resources/{id}").Handler(httptransport.NewServer(
		e.DeleteResourceEndpoint,
		decodeDeleteResourceRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindResourceByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindResourceByIDRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeFindResourcesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.FindResourcesRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.CreateResourceRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.UpdateResourceRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.DeleteResourceRequest
	if err := decodeHTTPRequest(r, &req); err != nil {
		return nil, err
	}
	return req, nil
}
