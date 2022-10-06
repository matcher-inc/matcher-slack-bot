package gss

import (
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func generateClientJWT() (*jwt.Config, error) {
	file, err := ioutil.ReadFile("config/spread_sheet_client.json")
	if err != nil {
		return nil, err
	}

	jwt, err := google.JWTConfigFromJSON(file, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return jwt, nil
}
