package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Send the message to Discord webhook
func SendDiscordWebhook(message string) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Printf("Error: DISCORD_WEBHOOK_URL is empty")
		return
	}

	payload := map[string]string{
		"content": message,
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

// Send the custom footer embed to Discord webhook
func SendFooterEmbed() {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Printf("Error: DISCORD_WEBHOOK_URL is empty")
		return
	}

	// Construct the embed object with footer information and color
	embed := map[string]interface{}{
		"footer": map[string]string{
			"text": "By RoMinjun - github.com/rominjun/yt-vid-discord-announcer",
		},
		"color": 0xFF0000, // YouTube red color
	}

	payload := map[string]interface{}{
		"embeds": []map[string]interface{}{embed},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Error sending custom message: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Error sending custom message: received status code %d", resp.StatusCode)
	}
}
