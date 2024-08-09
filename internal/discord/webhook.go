package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Send the custom embed message to Discord webhook
func SendDiscordWebhook(channelName, videoTitle, videoUrl string) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Printf("Error: DISCORD_WEBHOOK_URL is empty")
		return
	}

	// Construct the embed object with the necessary details
	embed := map[string]interface{}{
		"color": 0xFF0000, // YouTube red color
		"author": map[string]string{
			"name":     channelName,
			"icon_url": "https://avatar-resolver.vercel.app/youtube-avatar/q?url=" + videoUrl,
		},
		"title": videoTitle,
		"url":   videoUrl,
		"image": map[string]string{
			"url": "https://avatar-resolver.vercel.app/youtube-thumbnail/q?url=" + videoUrl,
		},
		"footer": map[string]interface{}{
			"text":     "RoMinjun | github.com/rominjun/yt-vid-discord-announcer",
			"icon_url": "https://avatars.githubusercontent.com/u/107768180?v=4",
		},
	}

	// Add the notification message with a clickable video title
	payload := map[string]interface{}{
		"content": "Hey @everyone, " + channelName + " just uploaded a new YouTube video! Go check it out!",
		"embeds":  []map[string]interface{}{embed},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Error sending webhook: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Error sending webhook: received status code %d", resp.StatusCode)
	}
}
