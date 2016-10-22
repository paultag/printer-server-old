package main

import (
	forecast "github.com/mlbright/forecast/v2"
)

type ForecastData struct {
	Forecast forecast.Forecast
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
