package gss

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type Sheet struct {
	service *sheets.Service
	sheetId string
}

func NewSheet(ctx context.Context, sheetId string) (*Sheet, error) {
	jwt, err := generateClientJWT()
	if err != nil {
		return nil, err
	}

	service, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}

	return &Sheet{
		service: service,
		sheetId: sheetId,
	}, nil
}
