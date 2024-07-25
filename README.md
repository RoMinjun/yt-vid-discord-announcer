# yt-vid-discord-announcer
A go project created to let Kirei's community know about a new youtube video

## .env file
Make sure you populate the .env file with the following:
```ini
CLIENT_ID=<google-oauth-client-id>
CLIENT_SECRET=<google-oauth-client-secret>
CHANNEL_ID=<channel_id>
REDIRECT_URL=http://localhost:8080/callback
DISCORD_WEBHOOK_URL=<discord-webhook>
```

## TODO:
- Add support for docker hosting