package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/openmesh/booking/endpoint"
)

func (s *Server) registerOrganizationRoutes(r *mux.Router) {
	e := endpoint.MakeOrganizationEndpoints(s.OrganizationService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET  /organization/  gets the organization of the currently authenticated user
	// PUT  /organization/  updates the organization of the currently authenticated user
	// POST /organizations/ creates a new organization

	r.Methods("GET").Path("/organization").Handler(httptransport.NewServer(
		e.FindCurrentOrganizationEndpoint,
		decodeFindCurrentOrganizationRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/organization").Handler(httptransport.NewServer(
		e.UpdateOrganizationEndpoint,
		decodeUpdateOrganizationRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("organizations").Handler(httptransport.NewServer(
		e.CreateOrganizationEndpoint,
		decodeCreateOrganizationRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindCurrentOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return empty{}, nil
}

func decodeUpdateOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Update); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Organization); err != nil {
		return nil, err
	}
	return req, nil
}
