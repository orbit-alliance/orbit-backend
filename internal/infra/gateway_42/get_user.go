package gateway_42

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUserByID(userID, token string) (userResponse, error) {
	var data userResponse
	url := fmt.Sprintf("https://api.intra.42.fr/v2/users/%s", userID)

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
