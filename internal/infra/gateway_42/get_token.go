package gateway_42

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func (g *Gateway42) getToken(ctx context.Context) (string, error) {
	id, secret := os.Getenv("CLIENT_ID_42"), os.Getenv("CLIENT_SECRET_42")
	if id == "" || secret == "" {
		return "", ErrSecretsIsNotDefined
	}

	/* -------------------------- cache in‑memory --------------------------- */
	g.mu.Lock()
	if g.token.AccessToken != "" && !g.token.expired() {
		tok := g.token.AccessToken
		g.mu.Unlock()
		return tok, nil
	}
	g.mu.Unlock()

	/* --------------------------- prepara chamada -------------------------- */
	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {id},
		"client_secret": {secret},
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(id+":"+secret))

	/* ----------------------------- faz POST ------------------------------- */
	t, err := doJSON[token42](
		ctx,
		g.client,
		http.MethodPost,
		"/oauth/token",
		auth, // cabeçalho Authorization completo
		form, // corpo x‑www‑form‑urlencoded
		ok2xx,
	)
	if err != nil {
		return "", fmt.Errorf("falha ao pedir token: %w", err)
	}
	if t.AccessToken == "" {
		return "", ErrNoAccessTokenInResponse
	}

	/* ---------------------------- grava cache ----------------------------- */
	g.mu.Lock()
	g.token = t
	g.token.CreatedAt = int(time.Now().Unix()) // garante coerência
	g.mu.Unlock()

	return t.AccessToken, nil
}
