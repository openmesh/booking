package oauth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/go-github/github"
	"github.com/openmesh/booking"
	"golang.org/x/oauth2"
	oauthgithub "golang.org/x/oauth2/github"
)

type OAuthService struct {
	authService booking.AuthService
	configs     map[string]*oauth2.Config
}

func NewOAuthService(authService booking.AuthService, cfgs map[string]*oauth2.Config) *OAuthService {
	return &OAuthService{
		authService: authService,
		configs:     cfgs,
	}
}

func (s *OAuthService) AddGitHub(clientID, clientSecret string) {
	s.configs[booking.AuthSourceGitHub] = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{},
		Endpoint:     oauthgithub.Endpoint,
	}
}

func (s *OAuthService) GetRedirectURL(
	ctx context.Context,
	req booking.GetRedirectURLRequest,
) booking.GetRedirectURLResponse {
	// Generate new OAuth state. Used to preevnt CSRF attacks.
	state := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, state)
	if err != nil {
		return booking.GetRedirectURLResponse{
			Err: fmt.Errorf("failed to read random byte into state: %w", err),
		}
	}
	// Encode state to string.
	encodedState := hex.EncodeToString(state)
	// Get cfg for specified authetication source.
	cfg, ok := s.configs[req.Source]
	if !ok {
		return booking.GetRedirectURLResponse{
			Err: booking.Errorf(
				booking.EAUTHSOURCENOTCONFIGURED,
				"Authentication source '%s' not configured",
				req.Source,
			),
		}
	}

	return booking.GetRedirectURLResponse{
		URL:   cfg.AuthCodeURL(encodedState),
		State: encodedState,
	}
}

func (s *OAuthService) HandleCallback(
	ctx context.Context,
	req booking.HandleCallbackRequest,
) booking.HandleCallbackResponse {
	switch req.Source {
	case booking.AuthSourceGitHub:

		userID, err := s.handleGitHubCallback(ctx, req.Code)
		if err != nil {
			return booking.HandleCallbackResponse{
				Err: fmt.Errorf("failed to handle github callback: %w", err),
			}
		}
		return booking.HandleCallbackResponse{
			UserID:      userID,
			RedirectURL: req.RedirectURL,
		}
	}
	return booking.HandleCallbackResponse{
		Err: booking.Errorf(
			booking.EAUTHSOURCEUNSUPPORTED,
			"Auth source '%s' is not supported",
			req.Source,
		),
	}
}

func (s *OAuthService) handleGitHubCallback(ctx context.Context, code string) (int, error) {
	cfg, ok := s.configs[booking.AuthSourceGitHub]
	if !ok {
		return 0, booking.Errorf(
			booking.EAUTHSOURCENOTCONFIGURED,
			"Authentication source '%s' not configured",
			booking.AuthSourceGitHub,
		)
	}
	tok, err := cfg.Exchange(ctx, code)
	if err != nil {
		return 0, fmt.Errorf("oauth exchange error: %w", err)
	}

	// Create a new GitHub API client.
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok.AccessToken},
	)))

	// Fetch user information for the currently authenticated user.
	// Require that we at least receive a user ID from GitHub.
	u, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return 0, fmt.Errorf("cannot fetch github user: %s", err)
	} else if u.ID == nil {
		return 0, fmt.Errorf("user ID not returned by GitHub, cannot authenticate user")
	}

	// Email is not necessarily available for all accounts. If it is, store it
	// so we can link together multiple OAuth providers in the future
	// (e.g. GitHub, Google, etc).
	var name string
	if u.Name != nil {
		name = *u.Name
	} else if u.Login != nil {
		name = *u.Login
	}
	var email string
	if u.Email != nil {
		email = *u.Email
	}
	var expiry *time.Time
	if !tok.Expiry.IsZero() {
		expiry = &tok.Expiry
	}

	// Create an authentication object with an associated user.
	auth := &booking.Auth{
		Source:       booking.AuthSourceGitHub,
		SourceID:     strconv.FormatInt(*u.ID, 10),
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		User: &booking.User{
			Name:  name,
			Email: email,
		},
	}
	if !tok.Expiry.IsZero() {
		auth.Expiry = &tok.Expiry
	}

	// Create the "Auth" object in the database. The AuthService will lookup
	// the user by email if they already exist. Otherwise, a new user will be
	// created and the user's ID will be set to auth.UserID.
	createAuthResp := s.authService.CreateAuth(ctx, booking.CreateAuthRequest{
		UserName:     name,
		UserEmail:    email,
		Source:       booking.AuthSourceGitHub,
		SourceID:     strconv.FormatInt(*u.ID, 10),
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       expiry,
	})
	if createAuthResp.Err != nil {
		return 0, fmt.Errorf("failed to create auth: %w", err)
	}

	return createAuthResp.Auth.UserID, nil
}
