package gateway_42

import "errors"

var (
	ErrSecretsIsNotDefined     = errors.New("Client ID or Client Secret is not defined")
	ErrNoAccessTokenInResponse = errors.New("No access token found in response")
)
