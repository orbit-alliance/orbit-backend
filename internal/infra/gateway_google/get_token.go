package gateway_google

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
)

func getToken() (string, error) {
	ctx := context.Background()
	var serviceFilename string = "seuarquivo-service-account.json"

	data, err := ioutil.ReadFile(serviceFilename)
	if err != nil {
		return "", ErrToReadFile
	}

	scopes := []string{
		"https://www.googleapis.com/auth/spreadsheets",
	}

	conf, err := google.JWTConfigFromJSON(data, scopes...)
	if err != nil {
		log.Fatalf("Erro ao criar JWT config: %v", err)
	}

	tokenSource := conf.TokenSource(ctx)
	token, err := tokenSource.Token()
	if err != nil {
		return "", ErrToGetToken
	}

	return token.AccessToken, nil
}
