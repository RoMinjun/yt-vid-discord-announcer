package service

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	YouTubeService *youtube.Service
	YouTubeMu      sync.Mutex
	tokenSource    oauth2.TokenSource
)

func InitializeYouTubeService(ctx context.Context, ts oauth2.TokenSource) error {
	YouTubeMu.Lock()
	defer YouTubeMu.Unlock()

	var err error
	YouTubeService, err = youtube.NewService(ctx, option.WithTokenSource(ts))
	tokenSource = ts
	return err
}

func ForceRefreshToken() (*oauth2.Token, error) {
	YouTubeMu.Lock()
	defer YouTubeMu.Unlock()

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %v", err)
	}

	return newToken, nil
}
