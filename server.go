package main

import (
	"fmt"
	"log"
	"os"

	"time"

	"encoding/json"
	"html/template"
	"net/http"

	"pault.ag/go/config"

	forecast "github.com/mlbright/forecast/v2"
	"pault.ag/go/politico"
)

type Cards struct {
	Lat           string `flag:"lat"            description:"Weather Latitude"`
	Lon           string `flag:"lon"            description:"Weather Longitude"`
	DarkskyAPIKey string `flag:"darksky-apikey" description:"Darksky API Key"`
}

type Page struct {
	Forecast forecast.Forecast
	News     []politico.Story
}

func newPhonyPage(cards Cards) (*Page, error) {
	fd, err := os.Open("/home/paultag/phony.json")
	if err != nil {
		return nil, err
	}
	page := Page{}
	return &page, json.NewDecoder(fd).Decode(&page)
}

func newPage(cards Cards) (*Page, error) {
	page := Page{}
	var err error

	f, err := forecast.Get(cards.DarkskyAPIKey, cards.Lat, cards.Lon, "now", forecast.US)
	if err != nil {
		return nil, err
	}

	page.Forecast = *f
	page.News, err = politico.News()
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func main() {
	conf := Cards{Lat: "38.897389", Lon: "-77.037410"}

	flags, err := config.LoadFlags("card", &conf)
	if err != nil {
		panic(err)
	}
	flags.Parse(os.Args[1:])

	fs := http.FileServer(http.Dir("output"))
	http.Handle("/output/", http.StripPrefix("/output/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//page, err := newPhonyPage(conf)
		page, err := newPage(conf)
		if err != nil {
			log.Printf("%s\n", err)
			return
		}
		index, err := template.New("").Funcs(template.FuncMap{
			"tempToString": func(measure float64) string {
				return fmt.Sprintf("%d", int(measure))
			},
			"importantWeather": func(slice []forecast.DataPoint) []*forecast.DataPoint {
				ret := []*forecast.DataPoint{}

				var lastState *forecast.DataPoint = nil
				for _, data := range slice {
					if lastState == nil {
						d := data
						lastState = &d
						continue
					}
					if data.Summary == lastState.Summary {
						continue
					}
					ret = append(ret, lastState)
					d := data
					lastState = &d
				}
				ret = append(ret, lastState)
				return ret
			},
			"topStories": func(on []politico.Story) []politico.Story {
				return on[:5]
			},
			"dateToString": func(when float64) string {
				return time.Unix(int64(when), 0).Format("03:04 PM")
			},
			"iconToFont": func(what string) string {
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
				}[what]
			},
		}).ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		if err := index.ExecuteTemplate(w, "index.html", &page); err != nil {
			log.Printf("%s\n", err)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
