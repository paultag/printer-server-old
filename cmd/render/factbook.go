package main

import (
	"encoding/json"
	"net/http"

	"pault.ag/go/printer-server/factbook"
)

type Factbook struct {
	URL string
}

func (f Factbook) Config() CardConfig {
	return CardConfig{Template: "factbook"}
}

func (f Factbook) Query() (interface{}, error) {
	resp, err := http.Get(f.URL)
	if err != nil {
		return nil, err
	}

	fb := factbook.Factbook{}

	if err := json.NewDecoder(resp.Body).Decode(&fb); err != nil {
		return nil, err
	}

	return &fb, nil
}

func NewFactbook(url string) (*Factbook, error) {
	return &Factbook{URL: url}, nil
}
