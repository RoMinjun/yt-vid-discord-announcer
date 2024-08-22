# yt-vid-discord-announcer
A go project created to let Kirei's community know about a new youtube video (and shorts). I've done this project as hobby so I'd get to know golang a little bit more with the goal to use it in my profession.
Go checkout [Kirei's youtube channel](https://www.youtube.com/@KireiLoL)

![](https://i.imgur.com/JDixW90.png)

# Getting started
Clone this repo on your host of choice:
```bash
git clone https://github.com/RoMinjun/yt-vid-discord-announcer.git && cd yt-vid-discord-announcer
```

## Adding thje .env file
Create the .env file with your selected language
<details>
  
<summary>Shell</summary>

```bash
touch .env
```

</details>

<details>
  
<summary>PowerShell</summary>

```powershell
New-Item -Name .env -ItemType File
```

</details>

Make sure you populate the .env file with the following using your favorite editor (must be `vi`):
> [!TIP] 
> I'd recommend using a domain that's publicly acessible so the callback will work when you're hosting it on docker, if not, localhost will do the trick.
```env
CLIENT_ID=<google-oauth-client-id>
CLIENT_SECRET=<google-oauth-client-secret>
CHANNEL_ID=<yt_channel_id>
REDIRECT_URL=http://localhost:8080/callback 
DISCORD_WEBHOOK_URL=<discord-webhook>
```

<br>

## Using docker

### building docker-compose (when using docker)
After you've populated the .env file in the root directory of the project, you can start the build:
```bash
docker compose up -d --build
```

and check the logging if desired to check whether the container is actually running or not:
```bash
docker compose logs -f
```


## Without docker
> [!IMPORTANT] 
> Recommended to use docker, but if you want it to work without docker add an absolute path to the .env

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

<br>

## Start checking for videos
To initially automate the video checking process, you need to login with your google account to your project.

Visit your host + port (http://localhost:8080 .e.g.) and login with your google account and consent to the project you've created for your google oauth configuration by clicking on `Advanced` and go to `localhost (unsafe)` and `continue`.
If you've done all of that, you'll be redirected to the `REDIRECT_URL` set in the `.env` file with the message `Authorization successful, you can close this tab.`


Now every 15 minutes the project will check for a new youtube video for the channel specified in the `.env`'s `CHANNEL_ID`. I wouldn't recommend decreasing the check time since I'm currently reaching 9.9/10k with 15 minutes, the sweet spot.
<br>

Unfortunately the oauth tokens expire after an hour, hence i've made a token refresher in the project. Each 55th minute it will use the refresh token to retrieve a new oauth-token and so on. Example from logging:
```bash
app-1  | 2024/08/15 19:20:51 Refreshed token: Access Token: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
app-1  | 2024/08/15 19:20:51 Token refreshed successfully
```





