package youtube

import (
	"html"
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

const lastNotifiedFilePath = "last_notified_id.txt"

// Save the last notified video ID to a file
func saveLastNotifiedId() {
	err := os.WriteFile(lastNotifiedFilePath, []byte(lastNotifiedId), 0644)
	if err != nil {
		log.Printf("Error saving last notified ID: %v", err)
	}
}

// Load the last notified video ID from a file
func loadLastNotifiedId() {
	data, err := os.ReadFile(lastNotifiedFilePath)
	if err != nil {
		log.Printf("Error loading last notified ID: %v", err)
		return
	}
	lastNotifiedId = string(data)
}

// StartCheckingYouTube initiates the process of checking YouTube for new videos
func StartCheckingYouTube() {
	loadLastNotifiedId()

	go func() {
		for {
			time.Sleep(55 * time.Minute)
			_, err := service.ForceRefreshToken()
			if err != nil {
				log.Printf("Error refreshing token: %v", err)
			} else {
				log.Println("Token refreshed successfully")
				// Immediately check for new videos after refreshing the token
				checkForNewVideos()
			}
		}
	}()

	for {
		checkForNewVideos()
		time.Sleep(15 * time.Minute) // Check every 15 minutes
	}
}

// checkForNewVideos checks YouTube for new videos and sends notifications
func checkForNewVideos() {
	checkMu.Lock()
	defer checkMu.Unlock()
	newVideoId, videoTitle, channelTitle := checkForNewVideo()
	if newVideoId != "" && newVideoId != lastNotifiedId {
		lastNotifiedId = newVideoId
		saveLastNotifiedId()
		videoLink := "https://www.youtube.com/watch?v=" + newVideoId
		message := "Hey @everyone," + channelTitle + " just uploaded [" + videoTitle + "](" + videoLink + ")! Go check it out!"
		discord.SendDiscordWebhook(message)
		log.Printf("New video found! URL: %s\n", videoLink)
	}
}

// checkForNewVideo fetches the most recent video from the YouTube channel
func checkForNewVideo() (string, string, string) {
	service.YouTubeMu.Lock()
	defer service.YouTubeMu.Unlock()

	channelId := os.Getenv("CHANNEL_ID")
	if channelId == "" {
		log.Printf("Error: CHANNEL_ID is empty")
		return "", "", ""
	}

	call := service.YouTubeService.Search.List([]string{"snippet"}).ChannelId(channelId).Order("date").PublishedAfter(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		if gErr, ok := err.(*googleapi.Error); ok {
			log.Printf("Google API error: %v", gErr)
		} else {
			log.Printf("Error checking for new video: %v", err)
		}
		return "", "", ""
	}

	if len(response.Items) == 0 {
		return "", "", ""
	}

	video := response.Items[0]
	videoTitle := html.UnescapeString(video.Snippet.Title) // Unescape HTML characters in the title
	return video.Id.VideoId, videoTitle, video.Snippet.ChannelTitle
}
