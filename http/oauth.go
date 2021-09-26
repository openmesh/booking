package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/openmesh/booking"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/openmesh/booking/endpoint"
)

func (s *Server) registerOAuthRoutes(r *mux.Router) {
	e := endpoint.MakeOAuthEndpoints(s.OAuthService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(addRequestToContext),
	}

	r.Methods("GET").Path("/oauth/{source}").Handler(httptransport.NewServer(
		e.GetRedirectURLEndpoint,
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			vars := mux.Vars(r)
			source, ok := vars["source"]
			if !ok {
				return nil, ErrBadRouting
			}
			// TODO: Ensure that source is valid.

			// Read session from request's cookies.
			_, err := s.session(r)
			if err != nil {
				return nil, fmt.Errorf("failed to read session: %w", err)
			}
			return booking.GetRedirectURLRequest{
				Source: source,
			}, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			res := response.(booking.GetRedirectURLResponse)
			session := Session{
				State: res.State,
			}
			err := s.setSession(w, session)
			if err != nil {
				return fmt.Errorf("failed to set session: %w", err)
			}

			r := requestFromContext(ctx)
			http.Redirect(w, r, res.URL, http.StatusFound)
			return nil
		},
		options...,
	))

	r.Methods("GET").Path("/oauth/{source}/callback").Handler(httptransport.NewServer(
		e.HandleCallbackEndpoint,
		func(c context.Context, r *http.Request) (request interface{}, err error) {
			vars := mux.Vars(r)
			source, ok := vars["source"]
			if !ok {
				return nil, ErrBadRouting
			}
			// TODO: Ensure that source is valid.

			// Read form variables passed in from GitHub.
			state, code := r.FormValue("state"), r.FormValue("code")

			// Read session from request.
			session, err := s.session(r)
			if err != nil {
				return nil, fmt.Errorf("cannot read session: %s", err)
			}

			// Validate that state matches session state.
			if state != session.State {
				return nil, fmt.Errorf("oauth state mismatch")
			}

			return booking.HandleCallbackRequest{
				Source:      source,
				Code:        code,
				RedirectURL: session.RedirectURL,
			}, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			r := requestFromContext(ctx)
			res := response.(booking.HandleCallbackResponse)
			// TODO: Handle error
			session, err := s.session(r)
			if err != nil {
				return fmt.Errorf("cannot read session: %w", err)
			}
			// Restore redirect URL stored on login
			redirectURL := session.RedirectURL

			// Update browser session to store the user's ID and clear OAuth state.
			session.UserID = res.UserID
			session.State = ""
			err = s.setSession(w, session)
			if err != nil {
				return fmt.Errorf("cannot set session cookie: %w", err)
			}

			// Redirect to store URL or, if not available, to the dashboard.
			if redirectURL == "" {
				redirectURL = "/dashboard"
			}
			http.Redirect(w, r, redirectURL, http.StatusFound)
			return nil
		},
		options...,
	))
}

// func decodeHandleCallbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	return booking.HandleCallbackRequest{}, nil
// }

func addRequestToContext(ctx context.Context, r *http.Request) context.Context {
	ctx = newContextWithRequest(ctx, r)
	return ctx
}
