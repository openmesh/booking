package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/openmesh/booking"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// Client represents an HTTP client.
type Client struct {
	URL string
}

// NewClient returns a new instance of Client.
func NewClient(u string) *Client {
	return &Client{URL: u}
}

// newRequest returns a new HTTP request but adds the current user's API key
// and sets the accept & content type headers to use JSON.
func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	// Build new request with base URL.
	req, err := http.NewRequest(method, c.URL+url, body)
	if err != nil {
		return nil, err
	}

	// Set API key in header.
	//if user := booking.UserFromContext(ctx); user != nil && user.APIKey != "" {
	//	req.Header.Set("Authorization", "Bearer "+user.APIKey)
	//}

	// Default to JSON format.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	return req, nil
}

// SessionCookieName is the name of the cookie used to store the session.
const SessionCookieName = "session"

// Session represents session data stored in a secure cookie.
type Session struct {
	UserID      int    `json:"userId"`
	RedirectURL string `json:"redirectUrl"`
	State       string `json:"state"`
}

// SetFlash sets the flash cookie for the next request to read.
func SetFlash(w http.ResponseWriter, s string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "flash",
		Value:    s,
		Path:     "/",
		HttpOnly: true,
	})
}

// Error prints & optionally logs an error message.
func Error(w http.ResponseWriter, r *http.Request, err error) {
	// Extract error code & message.
	code, message := booking.ErrorCode(err), booking.ErrorMessage(err)

	// Log & report internal errors.
	if code == booking.EINTERNAL {
		booking.ReportError(r.Context(), err, r)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(ErrorStatusCode(code))
	json.NewEncoder(w).Encode(&ErrorResponse{Error: message})
}

// ErrorResponse represents a JSON structure for error output.
type ErrorResponse struct {
	Error string `json:"error"`
}

// parseResponseError parses an JSON-formatted error response.
func parseResponseError(resp *http.Response) error {
	defer resp.Body.Close()

	// Read the response body, so we can reuse it for the error message if it
	// fails to decode as JSON.
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse JSON formatted error response.
	// If not JSON, use the response body as the error message.
	var errorResponse ErrorResponse
	if err := json.Unmarshal(buf, &errorResponse); err != nil {
		message := strings.TrimSpace(string(buf))
		if message == "" {
			message = "Empty response from server."
		}
		return booking.Errorf(FromErrorStatusCode(resp.StatusCode), message)
	}
	return booking.Errorf(FromErrorStatusCode(resp.StatusCode), errorResponse.Error)
}

// lookup of application error codes to HTTP status codes.
var codes = map[string]int{
	booking.ECONFLICT:       http.StatusConflict,
	booking.EINVALID:        http.StatusBadRequest,
	booking.ENOTFOUND:       http.StatusNotFound,
	booking.ENOTIMPLEMENTED: http.StatusNotImplemented,
	booking.EUNAUTHORIZED:   http.StatusUnauthorized,
	booking.EINTERNAL:       http.StatusInternalServerError,
}

// ErrorStatusCode returns the associated HTTP status code for a booking error code.
func ErrorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusInternalServerError
}

// FromErrorStatusCode returns the associated booking code for an HTTP status code.
func FromErrorStatusCode(code int) string {
	for k, v := range codes {
		if v == code {
			return k
		}
	}
	return booking.EINTERNAL
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(booking.Errorer); ok && e.Error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.Error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	// If the error is an application error then get the status code and body from the error.
	// Otherwise, return a 500 error status code.
	if _, ok := err.(booking.Error); ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(ErrorStatusCode(err.(booking.Error).Code))
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type empty struct{}
