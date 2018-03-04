package main

import (
	"fmt"

	forecast "github.com/mlbright/forecast/v2"
)

type ForecastDataPoint struct {
	forecast.DataPoint
}

type ForecastData struct {
	Forecast forecast.Forecast
}

func (f ForecastData) High() int64 {
	return int64(f.Forecast.Daily.Data[0].TemperatureMax)
}

func (f ForecastData) Low() int64 {
	return int64(f.Forecast.Daily.Data[0].TemperatureMin)
}

func (f ForecastData) Icon() string {
	icon, ok := map[string]string{
		"clear-day":           "wi-day-sunny",
		"clear-night":         "wi-night-clear",
		"rain":                "wi-day-rain",
		"snow":                "wi-day-snow",
		"sleet":               "wi-day-sleet",
		"wind":                "wi-day-windy",
		"fog":                 "wi-day-fog",
		"cloudy":              "wi-day-cloudy",
		"partly-cloudy-day":   "wi-day-cloudy",
		"partly-cloudy-night": "wi-night-partly-cloudy",
	}[f.Forecast.Daily.Data[0].Icon]
	if !ok {
		return "wi-na"
	}
	return icon
}

func (f ForecastDataPoint) HumidityString() string {
	return fmt.Sprintf("%d%%", int(f.Humidity*100))
}

func (f ForecastDataPoint) TemperatureMaxString() string {
	return fmt.Sprintf("%d°", int(f.TemperatureMax))
}

func (f ForecastDataPoint) TemperatureMinString() string {
	return fmt.Sprintf("%d°", int(f.TemperatureMin))
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

type Forecast struct {
	lat    string
	lon    string
	apikey string
}

func (f Forecast) Config() CardConfig {
	return CardConfig{Template: "weather"}
}

func (f Forecast) Query() (interface{}, error) {
	data, err := forecast.Get(f.apikey, f.lat, f.lon, "now", forecast.US, forecast.English)
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
