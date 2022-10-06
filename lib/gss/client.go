package gss

import (
	"context"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	service *sheets.Service
	sheetId string
}

func NewClient(ctx context.Context, sheetId string) (*Client, error) {
	file, err := ioutil.ReadFile("config/spread_sheet_client.json")
	if err != nil {
		return nil, err
	}

	jwt, err := google.JWTConfigFromJSON(file, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	service, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}

	return &Client{
		service: service,
		sheetId: sheetId,
	}, nil
}
