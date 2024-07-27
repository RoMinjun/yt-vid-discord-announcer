package tokenstore

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

func SaveToken(token *oauth2.Token) error {
	f, err := os.Create("token.json")
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func LoadToken() (*oauth2.Token, error) {
	f, err := os.Open("token.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
