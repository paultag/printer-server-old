package main

import (
	"fmt"
	"time"

	forecast "github.com/mlbright/forecast/v2"
)

type ForecastDataPoint struct {
	forecast.DataPoint
}

type ForecastData struct {
	Forecast forecast.Forecast
}

func (f ForecastDataPoint) HumidityString() string {
	return fmt.Sprintf("%d%%", int(f.Humidity*100))
}

func (f ForecastDataPoint) WeatherIcon() string {
	return map[string]string{
		"clear-day":           "wi-day-sunny",
		"clear-night":         "wi-night-clear",
		"rain":                "wi-day-rain",
		"snow":                "wi-day-snow",
		"sleet":               "wi-day-sleet",
		"wind":                "wi-day-windy",
		"fog":                 "wi-fog",
		"cloudy":              "wi-day-cloudy",
		"partly-cloudy-day":   "wi-day-cloudy",
		"partly-cloudy-night": "wi-night-cloudy",
	}[f.Icon]
}

func (f ForecastData) Today() ForecastDataPoint {
	return ForecastDataPoint{f.Forecast.Daily.Data[0]}
}

func (f ForecastData) DailyBarchart() Barchart {
	bars := []Bar{}
	height := 130

	for index, el := range f.Forecast.Hourly.Data[:22] {
		barHeight := int(el.Temperature)

		bar := Bar{
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
		}

		if index%2 == 0 {
			bar.YLabel = Label{
				X:      (index * 21),
				Y:      height + 7,
				Text:   time.Unix(int64(el.Time), 0).Format("03:04 PM"),
				Rotate: 45,
			}
		}

		bars = append(bars, bar)
	}

	return Barchart{
		Width:  472,
		Height: height + 60,
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
