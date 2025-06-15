package gateway_42

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func getUserByID(user42ID, token string) (userResponse, error) {
	var data userResponse
	url := fmt.Sprintf("https://api.intra.42.fr/v2/users/%s", user42ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return data, fmt.Errorf("error fetching token: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return data, nil
}

func (g *Gateway42) GetBasicUserInfo(
	ctx context.Context,
	token string,
) (*user.UserBasicInfoDTO, error) {

	// shape exato que a API devolve
	type userMe struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	}

	u, err := doJSON[userMe](
		ctx,
		g.client,
		http.MethodGet,
		"/v2/me",
		token,
		nil,
		ok2xx,
	)
	if err != nil {
		return nil, ErrFailToGetUserIn42
	}
	return &user.UserBasicInfoDTO{
		ID42:  fmt.Sprintf("%d", u.ID),
		Login: u.Login,
	}, nil
}
