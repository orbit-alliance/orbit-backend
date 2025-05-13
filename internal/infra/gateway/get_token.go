package gateway_42

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

func GetToken() (string, error) {
	clientID := os.Getenv("CLIENT_ID_42")
	clientSecret := os.Getenv("CLIENT_SECRET_42")

	if clientID == "" || clientSecret == "" {
		return "", ErrSecretsIsNotDefined
	}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)

	req, err := http.NewRequest("POST", "https://api.intra.42.fr/oauth/token", bytes.NewBufferString(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New("Error fetching token: " + resp.Status)
	}

	var data TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.AccessToken == "" {
		return "", ErrNoAccessTokenInResponse
	}

	return data.AccessToken, nil
}
