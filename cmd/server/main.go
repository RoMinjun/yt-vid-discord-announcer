package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rominjun/yt-vid-discord-announcer/internal/auth"
)

func main() {
	// Get the absolute path to the .env file
	// Adjust the path as necessary based on your directory structure
	envPath, err := filepath.Abs("../../.env")
	if err != nil {
		log.Fatalf("Error getting absolute path to .env file: %v", err)
	}

	// Log the absolute path to ensure it is correct
	log.Printf("Loading .env file from: %s", envPath)

	// Load the .env file
	err = godotenv.Load(envPath)
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
