package main

import (
	"fmt"

	forecast "github.com/mlbright/forecast/v2"
)

type ForecastData struct {
	Forecast forecast.Forecast
}

type Label struct {
	X    int
	Y    int
	Text string
}

type Bar struct {
	Width  int
	Height int
	X      int
	Y      int
	Label  Label
}

type Barchart struct {
	Width  int
	Height int
	Bars   []Bar
}

func (f ForecastData) DailyBarchart() Barchart {
	bars := []Bar{}
	height := 200

	for index, el := range f.Forecast.Hourly.Data[:22] {
		barHeight := int(el.Temperature)
		bars = append(bars, Bar{
			Width:  20,
			Height: barHeight,
			X:      (index * 21),
			Y:      (height - barHeight),
			Label: Label{
				X:    (index * 20) + 10,
				Y:    (height - barHeight) - 5,
				Text: fmt.Sprintf("%d°", int(el.Temperature)),
			},
		})
	}

	return Barchart{
		Width:  472,
		Height: height,
		Bars:   bars,
	}
}

type Forecast struct {
	lat    string
	lon    string
	apikey string
}

func (f Forecast) Config() CardConfig {
	return CardConfig{Template: "weather"}
}

func (f Forecast) Query() (interface{}, error) {
	data, err := forecast.Get(f.apikey, f.lat, f.lon, "now", forecast.US)
	return &ForecastData{
		Forecast: *data,
	}, err
}

func NewForecast(apikey, lat, lon string) (*Forecast, error) {
	return &Forecast{
		lat:    lat,
		lon:    lon,
		apikey: apikey,
	}, nil
}
