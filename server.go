package main

import (
	"encoding/json"
	"os"

	"time"

	"html/template"
	"net/http"

	forecast "github.com/mlbright/forecast/v2"
)

type Page struct {
	Forecast forecast.Forecast
}

func newPage() Page {
	fd, err := os.Open("/home/paultag/darksky.json")
	if err != nil {
		panic(err)
	}
	page := Page{}
	if err := json.NewDecoder(fd).Decode(&page.Forecast); err != nil {
		panic(err)
	}
	return page
}

func main() {
	index, err := template.New("").Funcs(template.FuncMap{
		"dateToString": func(when float64) string {
			return time.Unix(int64(when), 0).Format("03:04 PM")
		},
		"iconToFont": func(what string) string {
			return map[string]string{
				"rain":   "wi-day-rain",
				"cloudy": "wi-day-cloudy",
				"clear":  "wi-day-sunny",
			}[what]
		},
	}).ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("output"))
	http.Handle("/output/", http.StripPrefix("/output/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := newPage()
		index.ExecuteTemplate(w, "index.html", &page)
	})
	http.ListenAndServe(":8080", nil)
}
