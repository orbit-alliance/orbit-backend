package gateway_google

import (
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func GetRetroativePresencesBy42ID(ID42 string) ([]user.UserPresenceInEventDTO, error) {
	token, err := getToken()
	if err != nil {
		return nil, ErrToGetToken
	}
	var headerAjustment int = 2
	var userId int = -1
	users, err := getAllUsers(token)

	if err != nil {
		return nil, ErrGettingAllUsers
	}
	for i, row := range users.Values {
		if len(row) > 0 && row[0] == ID42 {
			userId = i + headerAjustment
			break
		}
	}
	if userId == -1 {
		return nil, ErrUserNotFound
	}
	eventData, err := getAllEvents(token)
	if err != nil {
		return nil, ErrGettingAllEvents
	}
	presences, err := getUserPresences(token, int64(userId))
	if err != nil {
		return nil, ErrGettingUserPresences
	}
	var userPresence []user.UserPresenceInEventDTO
	for i, _ := range eventData.Values[0] {
		if i < len(presences.Values) && len(presences.Values[i]) > 0 {
			var events []user.UserPresenceInEventDTO = make([]user.UserPresenceInEventDTO, 0, len(eventData.Values[0]))
			events = append(events, user.UserPresenceInEventDTO{
				Name:        eventData.Values[0][i],
				Date:        eventData.Values[1][i],
				HasPresence: presences.Values[0][i] == "TRUE",
			})
			userPresence = events
		}
	}
	return userPresence, nil
}
