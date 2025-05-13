package gateway_42

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserByID(userID, token string) (UserResponse, error) {
	var data UserResponse
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

	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
