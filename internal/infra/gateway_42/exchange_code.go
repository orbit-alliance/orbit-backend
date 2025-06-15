package gateway_42

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
)

func (g *Gateway42) ExchangeCodeForToken(
	ctx context.Context,
	code string,
) (string, error) {

	clientID := os.Getenv("CLIENT_ID_42")
	clientSecret := os.Getenv("CLIENT_SECRET_42")
	redirectURI := os.Getenv("REDIRECT_URI_42")

	form := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {redirectURI},
	}

	// Header Authorization: Basic base64(clientID:clientSecret)
	cred := clientID + ":" + clientSecret
	authHd := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred))

	type tokenResp struct {
		AccessToken string `json:"access_token"`
	}

	tr, err := doJSON[tokenResp](
		ctx,
		g.client, // reaproveita timeout, retry, etc.
		http.MethodPost,
		"/oauth/token",
		authHd, // valor COMPLETO do Authorization
		form,   // corpo x-www-form-urlencoded
		ok2xx,
	)
	if err != nil {
		return "", err
	}
	if tr.AccessToken == "" {
		return "", ErrNoAccessTokenInResponse
	}
	return tr.AccessToken, nil
}
