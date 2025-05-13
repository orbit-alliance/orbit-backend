package gateway_42

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLocationByUserID(startAt, endAt, userID, token string) (LocationResponse, error) {
	url := fmt.Sprintf("https://api.intra.42.fr/v2/users/%s/locations_stats?begin_at=%s&end_at=%s", userID, startAt, endAt)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error fetching token: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data LocationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
