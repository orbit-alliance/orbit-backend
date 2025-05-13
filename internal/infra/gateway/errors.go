package gateway_42

import "errors"

var (
	ErrSecretsIsNotDefined     = errors.New("Client ID or Client Secret is not defined")
	ErrNoAccessTokenInResponse = errors.New("No access token found in response")
	ErrFailToCreate42Token     = errors.New("user: fail to create 42 token")
	ErrFailToGetLocationIn42   = errors.New("user: fail to get location in 42")
	ErrFailToGetUserIn42       = errors.New("user: fail to get user in 42")
)
