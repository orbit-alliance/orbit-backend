package gateway_42

import (
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
)

type Client struct {
	http        *http.Client // permite injetar mock em testes
	baseURL     string       // "https://api.intra.42.fr"
	retryPolicy backoff.BackOff
}

// construtor padrão com timeout de 3 s e back‑off exponencial + jitter
func NewClient(timeout time.Duration, maxElapsed time.Duration) *Client {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = maxElapsed

	return &Client{
		http: &http.Client{
			Timeout: timeout,
		},
		baseURL:     "https://api.intra.42.fr",
		retryPolicy: bo,
	}
}
