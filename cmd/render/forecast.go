package main

import (
	"fmt"
	"time"

	forecast "github.com/mlbright/forecast/v2"
)

type ForecastData struct {
	Forecast forecast.Forecast
}

func (f ForecastData) DailyBarchart() Barchart {
	bars := []Bar{}
	height := 130

	for index, el := range f.Forecast.Hourly.Data[:22] {
		barHeight := int(el.Temperature)

		bars = append(bars, Bar{
			Width:  20,
			Height: barHeight,
			X:      (index * 21),
			Y:      (height - barHeight),
			Label: Label{
				X:      (index * 21) + 10,
				Y:      (height - barHeight) - 5,
				Text:   fmt.Sprintf("%d°", int(el.Temperature)),
				Rotate: -45,
			},
			YLabel: Label{
				X:      (index * 21),
				Y:      height + 7,
				Text:   time.Unix(int64(el.Time), 0).Format("03:04 PM"),
				Rotate: 45,
			},
		})
	}

	return Barchart{
		Width:  472,
		Height: height + 50,
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
