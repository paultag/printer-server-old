package main

import (
	"encoding/json"
	"os"
	"time"
	// "net/http"

	"pault.ag/go/printer-server/factbook"
)

type Factbook struct {
	URL string
}

type Response struct {
	Country  factbook.Country
	Template string
}

func (f Factbook) Config() CardConfig {
	return CardConfig{Template: "factbook"}
}

func (f Factbook) Query() (interface{}, error) {
	when := time.Now()

	body, err := os.Open(f.URL)
	// resp, err := http.Get(f.URL)
	if err != nil {
		return nil, err
	}

	fb := factbook.Factbook{}

	if err := json.NewDecoder(body).Decode(&fb); err != nil {
		// if err := json.NewDecoder(resp.Body).Decode(&fb); err != nil {
		return nil, err
	}

	/* ISO Weeks start on Monday */
	templates := []string{
		"transnational-issues",
		"overview",
		"economy",
		"terrorism",
		"government-branches",
		"illicit-drugs",
		"etemology",
	}

	return Response{
		Country:  fb.CountryOfTheWeek(when),
		Template: templates[when.Weekday()],
	}, nil
}

func NewFactbook(url string) (*Factbook, error) {
	return &Factbook{
		URL: url,
	}, nil
}
