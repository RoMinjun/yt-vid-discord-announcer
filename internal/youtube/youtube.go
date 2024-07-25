package youtube

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rominjun/yt-vid-discord-announcer/internal/discord"
	"github.com/rominjun/yt-vid-discord-announcer/internal/service"
	"google.golang.org/api/googleapi"
)

var (
	checkMu        sync.Mutex
	lastNotifiedId string
)

func StartCheckingYouTube() {
	go func() {
		for {
			time.Sleep(55 * time.Minute)
			_, err := service.ForceRefreshToken()
			if err != nil {
				log.Printf("Error refreshing token: %v", err)
			} else {
				log.Println("Token refreshed successfully")
			}
		}
	}()

	for {
		checkMu.Lock()
		newVideoId, videoTitle := checkForNewVideo()
		if newVideoId != "" && newVideoId != lastNotifiedId {
			lastNotifiedId = newVideoId
			videoLink := "https://www.youtube.com/watch?v=" + newVideoId
			message := "Hey @everyone, youtuber just uploaded [" + videoTitle + "](" + videoLink + ")"
			discord.SendDiscordWebhook(message)
			fmt.Printf("New video found! URL: %s\n", videoLink)
		}
		checkMu.Unlock()
		time.Sleep(15 * time.Minute) // Check every 15 minutes
	}
}

func checkForNewVideo() (string, string) {
	service.YouTubeMu.Lock()
	defer service.YouTubeMu.Unlock()

	channelId := os.Getenv("CHANNEL_ID")
	if channelId == "" {
		log.Printf("Error: CHANNEL_ID is empty")
		return "", ""
	}

	call := service.YouTubeService.Search.List([]string{"snippet"}).ChannelId(channelId).Order("date").MaxResults(1)
	response, err := call.Do()
	if err != nil {
		if gErr, ok := err.(*googleapi.Error); ok {
			if gErr.Code == 403 {
				// Quota exceeded, backoff
				log.Println("Quota exceeded, backing off for 30 mins")
				time.Sleep(30 * time.Minute)
				return "", ""
			}
		}
		log.Printf("Error making API call: %v", err)
		return "", ""
	}

	if len(response.Items) == 0 {
		log.Println("No videos found")
		return "", ""
	}

	video := response.Items[0]
	videoId := video.Id.VideoId
	videoTitle := video.Snippet.Title
	return videoId, videoTitle
}
