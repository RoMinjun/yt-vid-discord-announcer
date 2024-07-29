# yt-vid-discord-announcer
A go project created to let Kirei's community know about a new youtube video

> [!WARNING] 
> Only works with docker, if you want it to work without docker add an absolute path to the .env


## .env file
Make sure you populate the .env file with the following:
```ini
CLIENT_ID=<google-oauth-client-id>
CLIENT_SECRET=<google-oauth-client-secret>
CHANNEL_ID=<channel_id>
REDIRECT_URL=http://localhost:8080/callback
DISCORD_WEBHOOK_URL=<discord-webhook>
```
