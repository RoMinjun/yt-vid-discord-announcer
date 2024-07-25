package auth

import (
	"context"
	"fmt"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
	oauthState  = "state-token"
	tokenSource oauth2.TokenSource
	tokenMu     sync.Mutex

	startedChecking   bool
	startedCheckingMu sync.Mutex
	ctx               = context.Background()
)

func SetupOAuthConfig() {
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.readonly"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_URL"),
	}
}

func ForceRefreshToken() (*oauth2.Token, error) {
	tokenMu.Lock()
	defer tokenMu.Unlock()

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token: %v", err)
	}

	newToken, err := oauthConfig.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %v", err)
	}

	tokenSource = oauthConfig.TokenSource(ctx, newToken)

	return newToken, nil
}
