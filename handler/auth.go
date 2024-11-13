package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sneaktricks/sport-matchmaking-match-service/auth"
)

// Function to exchange authorization code for token
func exchangeCodeForToken(code string) (string, error) {
	tokenEndpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", auth.KeycloakURL, auth.Realm)
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:8080/callback")
	data.Set("client_id", auth.ClientID)
	data.Set("client_secret", auth.ClientSecret)

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}
