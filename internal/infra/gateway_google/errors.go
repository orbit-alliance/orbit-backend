package gateway_google

import "errors"

var (
	ErrSecretsIsNotDefined  = errors.New("Client ID or Client Secret is not defined")
	ErrGettingAllUsers      = errors.New("Error getting all users")
	ErrUserNotFound         = errors.New("User not found")
	ErrGettingUserPresences = errors.New("Error getting user presences")
	ErrGettingAllEvents     = errors.New("Error getting all events")
	ErrToGetToken           = errors.New("Error to get token")
	ErrToReadFile           = errors.New("Error to read file")
)
