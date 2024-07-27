package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rominjun/yt-vid-discord-announcer/internal/auth"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Log the value of DISCORD_WEBHOOK_URL to verify it's being loaded
	log.Printf("DISCORD_WEBHOOK_URL: %s", os.Getenv("DISCORD_WEBHOOK_URL"))
	log.Printf("CHANNEL_ID: %s", os.Getenv("CHANNEL_ID"))

	auth.SetupOAuthConfig()

	http.HandleFunc("/", auth.HandleMain)
	http.HandleFunc("/callback", auth.HandleCallback)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
