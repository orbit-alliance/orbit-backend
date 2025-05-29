package gateway_google

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sheetResponse struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

func getAllUsers(token string) (sheetResponse, error) {
	var data sheetResponse
	url := fmt.Sprintf("%s/%s/values/%s!B3:B", baseUrl, sheetId, presencePageName)

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

func getAllEvents(token string) (sheetResponse, error) {
	var data sheetResponse
	url := fmt.Sprintf("%s/%s/values/%s!E1:2", baseUrl, sheetId, presencePageName)

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

func getUserPresences(token string, userId int64) (sheetResponse, error) {
	var data sheetResponse
	url := fmt.Sprintf("%s/%s/values/%s!E%d:%d", baseUrl, sheetId, presencePageName, userId, userId)

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
