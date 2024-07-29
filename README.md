# yt-vid-discord-announcer
A go project created to let Kirei's community know about a new youtube video

> [!WARNING] 
> Only works with docker, if you want it to work without docker add an absolute path to the .env

Example without docker in `main.go`:
```go
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
err := godotenv.Load()
if err != nil {
  log.Fatalf("Error loading .env file: %v", err)
}
```


## .env file
Make sure you populate the .env file with the following:
```env
CLIENT_ID=<google-oauth-client-id>
CLIENT_SECRET=<google-oauth-client-secret>
CHANNEL_ID=<channel_id>
REDIRECT_URL=http://localhost:8080/callback
DISCORD_WEBHOOK_URL=<discord-webhook>
```
