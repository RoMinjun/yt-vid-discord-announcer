package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendDiscordWebhook(message string) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Printf("Error: DISCORD_WEBHOOK_URL is empty")
		return
	}

	payload := map[string]string{"content": message}
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
