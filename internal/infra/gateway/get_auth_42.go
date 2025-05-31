package gateway_42

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (g *Gateway42) GetBasicUserInfo(ctx context.Context, token string) (string, string, error) {
	url := "https://api.intra.42.fr/v2/me"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("erro na consulta ao /v2/me: %s", string(body))
	}

	var result struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprint(result.ID), result.Login, nil
}
