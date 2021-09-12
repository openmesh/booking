package http

import (
	"context"
	"encoding/json"
	"github.com/openmesh/booking"
	"net/http"
	"strconv"

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
		endpoint.MakeFindResourceByIDEndpoint(s.Handler),
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

func decodeFindResourceByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	return booking.FindResourceByIDQuery{ID: id}, nil
}

func decodeFindResourcesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.FindResourcesRequest{}

	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		req.Filter.ID = &id
	}

	name := r.URL.Query().Get("name")
	if name != "" {
		req.Filter.Name = &name
	}

	description := r.URL.Query().Get("description")
	if description != "" {
		req.Filter.Description = &description
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, err
		}
		req.Filter.Offset = offset
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, err
		}
		req.Filter.Limit = limit
	}

	orderBy := r.URL.Query().Get("orderBy")
	if orderBy != "" {
		req.Filter.OrderBy = &orderBy
	}

	return req, nil
}

func decodeCreateResourceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateResourceRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Resource); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateResourceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateResourceRequest
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	req.ID = id
	if err := json.NewDecoder(r.Body).Decode(&req.Update); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteResourceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return endpoint.DeleteResourceRequest{ID: id}, nil
}
