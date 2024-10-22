package http

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/openmesh/booking"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 1 * time.Second

// Server represents an HTTP server. It is meant to wrap all HTTP functionality
// used by the application so that dependent packages (such as cmd/wtfd) do not
// need to reference the "net/http" package at all.
type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router
	sc     *securecookie.SecureCookie
	logger log.Logger

	// Bind address & domain for the server's listener.
	// If domain is specified, server is run on TLS using acme/autocert.
	Addr   string
	Domain string

	// Keys used for secure cookie encryption.
	HashKey  string
	BlockKey string

	// GitHub OAuth settings.
	GitHubClientID     string
	GitHubClientSecret string

	// Services used by the various HTTP routes.
	AuthService           booking.AuthService
	AvailabilityService   booking.AvailabilityService
	BookingService        booking.BookingService
	EventService          booking.EventService
	OAuthService          booking.OAuthService
	OrganizationService   booking.OrganizationService
	ReportService         booking.ReportService
	ResourceService       booking.ResourceService
	TokenService          booking.TokenService
	UnavailabilityService booking.UnavailabilityService
	UserService           booking.UserService
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	// Create a new server that wraps the net/http server & add a gorilla router.
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// Report panics to external service.
	s.router.Use(reportPanic)

	// Our router is wrapped by another function handler to perform some
	// middleware-like tasks that cannot be performed by actual middleware.
	// This includes changing route paths for JSON endpoints & overridding methods.
	s.server.Handler = http.HandlerFunc(s.serveHTTP)

	// Setup error handling routes.
	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	// Setup endpoint to display deployed version.
	s.router.HandleFunc("/debug/version", s.handleVersion).Methods("GET")
	s.router.HandleFunc("/debug/commit", s.handleCommit).Methods("GET")

	// Setup a base router that excludes asset handling.
	router := s.router.PathPrefix("/").Subrouter()
	router.Use(loadFlash)

	// Handle embedded asset serving. This serves files embedded from http/assets.
	fs := http.FileServer(http.Dir("./http/assets"))
	s.router.PathPrefix("/static/").Handler(fs)
	router.HandleFunc("/manifest", s.handleManifest).Methods("GET")

	// router.Use(trackMetrics)

	// Handle authentication check within handler function for home page.
	// router.HandleFunc("/", s.handleIndex).Methods("GET")

	return s
}

func (s *Server) RegisterRoutes() {
	authSpaRoutes := []string{
		"/signup",
		"/dashboard",
	}

	noAuthSpaRoutes := []string{
		"/",
	}
	s.router.Use(s.authenticate)
	// Register unauthenticated routes.
	{
		r := s.router.PathPrefix("/").Subrouter()
		r.Use(s.requireNoAuth)
		s.registerOAuthRoutes(r)
		s.registerOrganizationRoutes(r)
		for _, route := range noAuthSpaRoutes {
			r.HandleFunc(route, s.handleSpaRoute).Methods("GET")
		}
	}
	// Register authenticated routes.
	{
		r := s.router.PathPrefix("/").Subrouter()
		r.Use(s.requireAuth)
		s.registerResourceRoutes(r)
		s.registerBookingRoutes(r)
		s.registerUnavailabilityRoutes(r)
		s.registerTokenRoutes(r)
		for _, route := range authSpaRoutes {
			r.HandleFunc(route, s.handleSpaRoute).Methods("GET")
		}
	}
}

func (s *Server) AttachLogger(logger log.Logger) {
	s.logger = logger
}

// UseTLS returns true if the cert & key file are specified.
func (s *Server) UseTLS() bool {
	return s.Domain != ""
}

// Scheme returns the URL scheme for the server.
func (s *Server) Scheme() string {
	if s.UseTLS() {
		return "https"
	}
	return "http"
}

// Port returns the TCP port for the running server.
// This is useful in tests where we allocate a random port by using ":0".
func (s *Server) Port() int {
	if s.ln == nil {
		return 0
	}
	return s.ln.Addr().(*net.TCPAddr).Port
}

// URL returns the local base URL of the running server.
func (s *Server) URL() string {
	scheme, port := s.Scheme(), s.Port()

	// Use localhost unless a domain is specified.
	domain := "localhost"
	if s.Domain != "" {
		domain = s.Domain
	}

	// Return without port if using standard ports.
	if (scheme == "http" && port == 80) || (scheme == "https" && port == 443) {
		return fmt.Sprintf("%s://%s", s.Scheme(), domain)
	}
	return fmt.Sprintf("%s://%s:%d", s.Scheme(), domain, s.Port())
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() (err error) {
	// Initialize our secure cookie with our encryption keys.
	if err := s.openSecureCookie(); err != nil {
		return err
	}

	// Validate GitHub OAuth settings.
	if s.GitHubClientID == "" {
		return fmt.Errorf("github client id required")
	} else if s.GitHubClientSecret == "" {
		return fmt.Errorf("github client secret required")
	}

	// Open a listener on our bind address.
	if s.Domain != "" {
		s.ln = autocert.NewListener(s.Domain)
	} else {
		if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
			return err
		}
	}

	// Begin serving requests on the listener. We use Serve() instead of
	// ListenAndServe() because it allows us to check for listen errors (such
	// as trying to use an already open port) synchronously.
	go s.server.Serve(s.ln)

	return nil
}

// openSecureCookie validates & decodes the block & hash key and initializes
// our secure cookie implementation.
func (s *Server) openSecureCookie() error {
	// Ensure hash & block key are set.
	if s.HashKey == "" {
		return fmt.Errorf("hash key required")
	} else if s.BlockKey == "" {
		return fmt.Errorf("block key required")
	}

	// Decode from hex to byte slices.
	hashKey, err := hex.DecodeString(s.HashKey)
	if err != nil {
		return fmt.Errorf("invalid hash key")
	}
	blockKey, err := hex.DecodeString(s.BlockKey)
	if err != nil {
		return fmt.Errorf("invalid block key")
	}

	// Initialize cookie management & encode our cookie data as JSON.
	s.sc = securecookie.New(hashKey, blockKey)
	s.sc.SetSerializer(securecookie.JSONEncoder{})

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// OAuth2Config returns the GitHub OAuth2 configuration.
func (s *Server) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     s.GitHubClientID,
		ClientSecret: s.GitHubClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// Override method for forms passing "_method" value.
	if r.Method == http.MethodPost {
		switch v := r.PostFormValue("_method"); v {
		case http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete:
			r.Method = v
		}
	}

	// Override content-type for certain extensions.
	// This allows us to easily cURL API endpoints with a ".json" or ".csv"
	// extension instead of having to explicitly set Content-type & Accept headers.
	// The extensions are removed so they don't appear in the routes.
	switch ext := path.Ext(r.URL.Path); ext {
	case ".json":
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Content-type", "application/json")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	case ".csv":
		r.Header.Set("Accept", "text/csv")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	}

	// Delegate remaining HTTP handling to the gorilla router.
	s.router.ServeHTTP(w, r)
}

// authenticate is middleware for loading session data from a cookie or API key header.
func (s *Server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Login via API key, if available.
		if v := r.Header.Get("Authorization"); strings.HasPrefix(v, "Bearer ") {
			apiKey := strings.TrimPrefix(v, "Bearer ")

			// Lookup organization by API key. Display error if not found.
			org, err := s.OrganizationService.FindOrganizationByPrivateKey(r.Context(), apiKey)
			if err != nil && booking.ErrorCode(err) == booking.ENOTFOUND {
				Error(w, r, booking.Errorf(booking.EUNAUTHORIZED, "Invalid API key."))
				return
			} else if err != nil {
				Error(w, r, err)
				return
			}

			// Update request context to include authenticated user.
			r = r.WithContext(booking.NewContextWithOrganization(r.Context(), org))

			// Delegate to next HTTP handler.
			next.ServeHTTP(w, r)
			return
		}

		// Read session from secure cookie.
		session, _ := s.session(r)

		// Read user, if available. Ignore if fetching assets.
		if session.UserID != 0 {
			if user, err := s.UserService.FindUserByID(r.Context(), session.UserID); err != nil {
				s.logger.Log("cannot find session user: id=%d err=%s", session.UserID, err)
			} else {
				r = r.WithContext(booking.NewContextWithUser(r.Context(), user))
				r = r.WithContext(booking.NewContextWithOrganization(r.Context(), user.Organization))
			}
		}

		next.ServeHTTP(w, r)
	})
}

// requireNoAuth is middleware for requiring no authentication.
// This is used if a user goes to log in but is already logged in.
func (s *Server) requireNoAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If user is logged in, redirect to the home page.
		if userID := booking.UserIDFromContext(r.Context()); userID != 0 {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(w, r)
	})
}

// requireAuth is middleware for requiring authentication. This is used by
// nearly every page except for the login & oauth pages.
func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if orgID := booking.OrganizationIDFromContext(r.Context()); orgID != 0 {
			next.ServeHTTP(w, r)
			return
		}
		// If user is logged in, delegate to next HTTP handler.
		if userID := booking.UserIDFromContext(r.Context()); userID != 0 {
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise save the current URL (without scheme/host).
		redirectURL := r.URL
		redirectURL.Scheme, redirectURL.Host = "", ""

		// Save the URL to the session and redirect to the log in page.
		// On successful login, the user will be redirected to their original location.
		session, _ := s.session(r)
		session.RedirectURL = redirectURL.String()
		if err := s.setSession(w, session); err != nil {
			s.logger.Log("http: cannot set session: %s", err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})
}

// loadFlash is middleware for reading flash data from the cookie.
// Data is only loaded once and then immediately cleared... hence the name "flash".
func loadFlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read & clear flash from cookies.
		if cookie, _ := r.Cookie("flash"); cookie != nil {
			SetFlash(w, "")
			r = r.WithContext(booking.NewContextWithFlash(r.Context(), cookie.Value))
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(w, r)
	})
}

// requestPathTemplate returns the route path template for r.
func requestPathTemplate(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route == nil {
		return ""
	}
	tmpl, _ := route.GetPathTemplate()
	return tmpl
}

// reportPanic is middleware for catching panics and reporting them.
func reportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				booking.ReportPanic(err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// handleNotFound handles requests to routes that don't exist.
func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not found"))
}

// handleVersion displays the deployed version.
func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(booking.Version))
}

// handleVersion displays the deployed commit.
func (s *Server) handleCommit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(booking.Commit))
}

// session returns session data from the secure cookie.
func (s *Server) session(r *http.Request) (Session, error) {
	// Read session data from cookie.
	// If it returns an error then simply return an empty session.
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return Session{}, nil
	}

	// Decode session data into a Session object & return.
	var session Session
	if err := s.UnmarshalSession(cookie.Value, &session); err != nil {
		return Session{}, err
	}
	return session, nil
}

// setSession creates a secure cookie with session data.
func (s *Server) setSession(w http.ResponseWriter, session Session) error {
	// Encode session data to JSON.
	buf, err := s.MarshalSession(session)
	if err != nil {
		return err
	}

	// Write cookie to HTTP response.
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    buf,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Secure:   s.UseTLS(),
		HttpOnly: true,
	})
	return nil
}

// MarshalSession encodes session data to string.
// This is exported to allow the unit tests to generate fake sessions.
func (s *Server) MarshalSession(session Session) (string, error) {
	return s.sc.Encode(SessionCookieName, session)
}

// UnmarshalSession decodes session data into a Session object.
// This is exported to allow the unit tests to generate fake sessions.
func (s *Server) UnmarshalSession(data string, session *Session) error {
	return s.sc.Decode(SessionCookieName, data, &session)
}

// ListenAndServeTLSRedirect runs an HTTP server on port 80 to redirect users
// to the TLS-enabled port 443 server.
func ListenAndServeTLSRedirect(domain string) error {
	return http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+domain, http.StatusFound)
	}))
}

// ListenAndServeDebug runs an HTTP server with /debug endpoints (e.g. pprof, vars).
func ListenAndServeDebug() error {
	h := http.NewServeMux()
	// h.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":6060", h)
}

func (s *Server) handleSpaRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./http/assets/index.html")
}

func (s *Server) handleManifest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./http/assets/manifest.json")
}
