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

func (s *Server) registerTokenRoutes(r *mux.Router) {
	e := endpoint.MakeTokenEndpoints(s.TokenService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.CreateTokenEndpoint,
		decodeCreateTokenRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.FindTokensEndpoint,
		decodeFindTokensRequest,
		encodeResponse,
		options...,
	))
}

func decodeCreateTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.CreateTokenRequest{}, nil
}

func decodeFindTokensRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.FindTokensRequest{}, nil
}
