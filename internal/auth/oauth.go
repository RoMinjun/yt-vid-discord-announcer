package auth

import (
	"context"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/rominjun/yt-vid-discord-announcer/internal/service"
	"github.com/rominjun/yt-vid-discord-announcer/internal/tokenstore"
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

	token, err := tokenstore.LoadToken()
	if err == nil {
		log.Printf("Loaded token: Access Token: %s, Refresh Token: %s", token.AccessToken, token.RefreshToken)
		tokenSource = oauthConfig.TokenSource(ctx, token)
		err = service.InitializeYouTubeService(ctx, oauthConfig, tokenSource)
		if err != nil {
			log.Fatalf("Error creating YouTube service: %v", err)
		}
	} else {
		log.Printf("No existing token found, user needs to authenticate %s", strings.TrimSuffix(os.Getenv("REDIRECT_URL"), "/callback"))
	}
}
