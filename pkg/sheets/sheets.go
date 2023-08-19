package sheets

import (
	"context"
	"errors"
	"io/fs"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type Spreadsheet struct {
	service       *sheets.Service
	spreadsheetId string
}

func NewSpreadsheet(f fs.ReadFileFS, spreadsheetId string) (Spreadsheet, error) {
	jsonKey, _ := f.ReadFile("credentials.json")

	conf, err := google.JWTConfigFromJSON(jsonKey, sheets.SpreadsheetsScope)
	if err != nil {
		return Spreadsheet{}, errors.New("sheets: can't get JWT from credentials.json")
	}

	client := conf.Client(context.TODO())
	service, err := sheets.New(client)
	if err != nil {
		return Spreadsheet{}, errors.New("sheets: can't create client for")
	}

	return Spreadsheet{service: service, spreadsheetId: spreadsheetId}, nil
}

func (s Spreadsheet) Get() (*sheets.ValueRange, error) {
	readRange := "A2:I200"
	resp, err := s.service.Spreadsheets.Values.Get(s.spreadsheetId, readRange).Do()

	if err != nil {
		return nil, errors.New("sheets: can't create client for")
	}

	return resp, nil
}
