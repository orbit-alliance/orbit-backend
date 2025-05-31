package gateway_42

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (g *Gateway42) ExchangeCodeForToken(ctx context.Context, code string) (string, error) {
	clientID := os.Getenv("CLIENT_ID_42")
	clientSecret := os.Getenv("CLIENT_SECRET_42")
	redirectURI := os.Getenv("REDIRECT_URI_42")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.intra.42.fr/oauth/token",
		strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Header Authorization → Basic base64(client_id:client_secret)
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro ao obter token: %s", string(body))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("token inválido")
	}

	return tokenResp.AccessToken, nil
}
