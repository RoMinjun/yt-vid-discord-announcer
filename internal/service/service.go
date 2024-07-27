package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/rominjun/yt-vid-discord-announcer/internal/tokenstore"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	YouTubeService *youtube.Service
	YouTubeMu      sync.Mutex
	tokenSource    oauth2.TokenSource
	oauthConfig    *oauth2.Config // Add this variable
)

func InitializeYouTubeService(ctx context.Context, config *oauth2.Config, ts oauth2.TokenSource) error {
	YouTubeMu.Lock()
	defer YouTubeMu.Unlock()

	var err error
	YouTubeService, err = youtube.NewService(ctx, option.WithTokenSource(ts))
	tokenSource = ts
	oauthConfig = config // Set the oauthConfig
	return err
}

func ForceRefreshToken() (*oauth2.Token, error) {
	YouTubeMu.Lock()
	defer YouTubeMu.Unlock()

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %v", err)
	}

	log.Printf("Refreshed token: Access Token: %s, Refresh Token: %s", newToken.AccessToken, newToken.RefreshToken)

	// Update the token source with the new token
	tokenSource = oauthConfig.TokenSource(context.Background(), newToken)

	// Save the new token
	err = tokenstore.SaveToken(newToken)
	if err != nil {
		return nil, fmt.Errorf("failed to save refreshed token: %v", err)
	}

	return newToken, nil
}
