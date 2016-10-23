package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"encoding/json"

	"html/template"
	"net/http"
	// "pault.ag/go/wmata"
)

func loadTemplates(root string) (*template.Template, error) {
	files := []string{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".html") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return template.New("").Funcs(
		map[string]interface{}{
			"humanDate": func(when time.Time) string {
				return when.Format("03:04 PM")
			},
		},
	).ParseFiles(files...)
}

type Server struct {
	cards []Card
}

func (s *Server) Add(c Card) {
	s.cards = append(s.cards, c)
}

func (s *Server) Query() (map[string]interface{}, error) {
	query := map[string]interface{}{}
	for _, card := range s.cards {
		config := card.Config()
		data, err := card.Query()
		if err != nil {
			return nil, err
		}
		query[config.Template] = data
	}

	return query, nil
}

func NewServer() (*Server, error) {
	return &Server{
		cards: []Card{},
	}, nil
}

type Config struct {
	SenseDirectory string
	WMATAAPIKey    string
	DarkSkyAPIKey  string
	Lat            string
	Lon            string
}

func loadConfig(path string) (*Config, error) {
	config := Config{}
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &config, json.NewDecoder(fd).Decode(&config)
}

func main() {
	server, err := NewServer()
	if err != nil {
		panic(err)
	}

	config, err := loadConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	// forecast, _ := NewForecast(
	// 	config.DarkSkyAPIKey,
	// 	config.Lat,
	// 	config.Lon,
	// )
	// server.Add(forecast)

	// politico, _ := NewPolitico()
	// server.Add(politico)

	// wmata, _ := NewWMATA(
	// 	config.WMATAAPIKey,
	// 	[]wmata.Line{wmata.GreenLine},
	// )
	// server.Add(wmata)

	sense, _ := NewSense(config.SenseDirectory)
	server.Add(sense)

	fs := http.FileServer(http.Dir("output"))
	http.Handle("/output/", http.StripPrefix("/output/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := loadTemplates("templates")
		if err != nil {
			log.Fatalf("%s\n", err)
			return
		}
		query, err := server.Query()
		if err != nil {
			log.Fatalf("%s\n", err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(200)
		if err := template.ExecuteTemplate(w, "index.html", query); err != nil {
			log.Println(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
