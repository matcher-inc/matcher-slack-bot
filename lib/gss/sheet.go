package gss

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type SpreadSheet struct {
	service       *sheets.Service
	spreadSheetId string
}

type Sheet struct {
	spreadSheet *SpreadSheet
	name        string
}

func GetSpreadSheetById(spreadSheetId string) (*SpreadSheet, error) {
	jwt, err := generateClientJWT()
	if err != nil {
		return nil, err
	}

	service, err := sheets.New(jwt.Client(context.Background()))
	if err != nil {
		return nil, err
	}

	return &SpreadSheet{
		service:       service,
		spreadSheetId: spreadSheetId,
	}, nil
}

func (ss SpreadSheet) GetSheetByName(name string) *Sheet {
	return &Sheet{
		spreadSheet: &ss,
		name:        name,
	}
}
