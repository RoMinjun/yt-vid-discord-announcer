package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rominjun/yt-vid-discord-announcer/internal/service"
	"github.com/rominjun/yt-vid-discord-announcer/internal/youtube"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthState {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthState, state)
		http.Error(w, "Invalid OAuth state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("oauthConfig.Exchange() failed with '%s'\n", err)
		http.Error(w, "OAuth exchange failed", http.StatusInternalServerError)
		return
	}

	tokenMu.Lock()
	defer tokenMu.Unlock()
	tokenSource = oauthConfig.TokenSource(ctx, token)
	err = service.InitializeYouTubeService(ctx, tokenSource)
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	fmt.Fprintf(w, "Authorization successful, you can close this tab.")

	startedCheckingMu.Lock()
	defer startedCheckingMu.Unlock()
	if !startedChecking {
		startedChecking = true
		go youtube.StartCheckingYouTube()
	}
}
