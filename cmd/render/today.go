package main

import (
	"time"
)

type Today struct{}

type TodayResponse struct {
	When time.Time
}

func (Today) Config() CardConfig {
	return CardConfig{Template: "today"}
}

func (Today) Query() (interface{}, error) {
	return TodayResponse{
		When: time.Now(),
	}, nil
}
