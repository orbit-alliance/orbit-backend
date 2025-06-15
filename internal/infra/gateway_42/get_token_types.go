package gateway_42

import "time"

type token42 struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	CreatedAt        int    `json:"created_at"`
	SecretValidUntil int    `json:"secret_valid_until"`
}

func (t token42) expired() bool {
	if t.AccessToken == "" || t.ExpiresIn == 0 {
		return true
	}
	exp := time.Unix(int64(t.CreatedAt), 0).Add(
		time.Duration(t.ExpiresIn-10) * time.Second,
	)
	return time.Now().After(exp)
}
