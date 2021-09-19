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
	errs := make([]booking.ValidationError, 0)
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errs = append(errs, booking.ValidationError{Name: "id", Reason: "Must be valid integer"})
	}
	resourceIdStr, ok := vars["resourceId"]
	if !ok {
		return nil, ErrBadRouting
	}
	resourceID, err := strconv.Atoi(resourceIdStr)
	if err != nil {
		errs = append(errs, booking.ValidationError{Name: "resourceId", Reason: "Must be valid integer"})
	}
	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return booking.FindUnavailabilityByIDRequest{ID: id, ResourceID: resourceID}, nil
}

func decodeFindUnavailabilitiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := booking.FindUnavailabilitiesRequest{}
	var errs []booking.ValidationError

	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{Name: "id", Reason: "Must be a valid integer"})
		}
		req.ID = &id
	}

	resourceIDStr := r.URL.Query().Get("resourceId")
	if resourceIDStr != "" {
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{Name: "resourceId", Reason: "Must be a valid integer"})
		}
		req.ResourceID = resourceID
	}

	if fromStr := r.URL.Query().Get("from"); fromStr != "" {
		from, err := booking.ParseTime(fromStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{Name: "from", Reason: "Unrecognized time format. Prefer RFC3339 formatting when submitting date times. https://datatracker.ietf.org/doc/html/rfc3339"})
		}
		req.From = &from
	}

	if toStr := r.URL.Query().Get("from"); toStr != "" {
		to, err := booking.ParseTime(toStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{Name: "to", Reason: "Unrecognized time format. Prefer RFC3339 formatting when submitting date times. https://datatracker.ietf.org/doc/html/rfc3339"})
		}
		req.To = &to
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

func decodeCreateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.CreateUnavailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	var errs []booking.ValidationError

	resourceIDStr := r.URL.Query().Get("resourceId")
	if resourceIDStr != "" {
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{Name: "resourceId", Reason: "Must be a valid integer"})
		}
		req.ResourceID = resourceID
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}

func decodeUpdateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booking.UpdateUnavailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
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

	resourceIDStr, ok := vars["resourceId"]
	if !ok {
		return nil, ErrBadRouting
	}
	if resourceIDStr != "" {
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "id",
				Reason: "Must be a valid integer",
			})
		}
		req.ResourceID = resourceID
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}

func decodeDeleteUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

	resourceIDStr, ok := vars["resourceId"]
	if !ok {
		return nil, ErrBadRouting
	}
	if resourceIDStr != "" {
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			errs = append(errs, booking.ValidationError{
				Name:   "resourceId",
				Reason: "Must be a valid integer",
			})
		}
		req.ResourceID = resourceID
	}

	if len(errs) > 0 {
		return nil, booking.WrapValidationErrors(errs)
	}

	return req, nil
}
