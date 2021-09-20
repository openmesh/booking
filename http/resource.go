package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, booking.WrapValidationErrors([]booking.ValidationError{{Name: "id", Reason: "Must be a valid integer"}})
	}
	return booking.FindResourceByIDRequest{ID: id}, nil
}

func decodeFindResourcesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := booking.FindResourcesRequest{}

	var errs []booking.ValidationError

	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "id",
				Reason: "Must be a valid integer",
			})
		}
		req.ID = &id
	}

	name := r.URL.Query().Get("name")
	if name != "" {
		req.Name = &name
	}

	description := r.URL.Query().Get("description")
	if description != "" {
		req.Description = &description
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "offset",
				Reason: "Must be a valid integer",
			})
		}
		req.Offset = offset
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "offset",
				Reason: "Must be a valid integer",
			})
		}
		req.Limit = limit
	}

	orderBy := r.URL.Query().Get("orderBy")
	if orderBy != "" {
		req.OrderBy = &orderBy
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}

func decodeCreateResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.CreateResourceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.UpdateResourceRequest
	vars := mux.Vars(r)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	var errs []booking.ValidationError
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "id",
				Reason: "Must be a valid integer",
			})
		}
		req.ID = id
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}

func decodeDeleteResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.DeleteUnavailabilityRequest
	var errs []booking.ValidationError
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "id",
				Reason: "Must be a valid integer",
			})
		}
		req.ID = id
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}
