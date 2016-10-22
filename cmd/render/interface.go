package main

type CardConfig struct {
	Template string // "index.html"
	// FuncMap?
}

type Card interface {
	Config() CardConfig
	Query() (interface{}, error)
}
