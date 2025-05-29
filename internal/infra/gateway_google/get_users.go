package gateway_google

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
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

func GetBy42ID(ID42 string) (user.UserPresenceInEventDTO, error) {
	token, err := getToken()
	if err != nil {
		return user.UserPresenceInEventDTO{}, ErrToGetToken
	}
	var headerAjustment int = 2
	var userId int = -1
	users, err := getAllUsers(token)

	if err != nil {
		return user.UserPresenceInEventDTO{}, ErrGettingAllUsers
	}
	for i, row := range users.Values {
		if len(row) > 0 && row[0] == ID42 {
			userId = i + headerAjustment
			break
		}
	}
	if userId == -1 {
		return user.UserPresenceInEventDTO{}, ErrUserNotFound
	}
	eventData, err := getAllEvents(token)
	if err != nil {
		return user.UserPresenceInEventDTO{}, ErrGettingAllEvents
	}
	presences, err := getUserPresences(token, int64(userId))
	if err != nil {
		return user.UserPresenceInEventDTO{}, ErrGettingUserPresences
	}
	var userPresence user.UserPresenceInEventDTO
	for i, _ := range eventData.Values[0] {
		if i < len(presences.Values) && len(presences.Values[i]) > 0 {
			userPresence.Events = append(userPresence.Events, struct {
				Name        string `json:"name"`
				Date        string `json:"date"`
				HasPresence bool   `json:"has_presence"`
			}{
				Name:        eventData.Values[0][i],
				Date:        eventData.Values[1][i],
				HasPresence: presences.Values[0][i] == "TRUE",
			})
		}
	}
	return userPresence, nil
}
