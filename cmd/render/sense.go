package main

import (
	"time"

	"pault.ag/go/sense"
)

type SenseData struct {
	Timeline *sense.Timeline
}

func (s SenseData) ImportantEvents() []sense.TimelineEvent {
	events := []sense.TimelineEvent{}

	for _, event := range s.Timeline.Events {
		if event.Type == "IN_BED" ||
			event.Type == "PARTNER_MOTION" ||
			event.Type == "GENERIC_MOTION" {
			continue
		}
		events = append(events, event)
	}

	return events
}

type Sense struct {
	senseAPI *sense.Sense
}

func (Sense) Config() CardConfig {
	return CardConfig{Template: "sense"}
}

func (s Sense) Query() (interface{}, error) {
	timeline, err := s.senseAPI.Timeline(
		time.Now().Add((-time.Hour * 24)).Format("2006-01-02"))
	if err != nil {
		panic(err)
	}

	return &SenseData{
		Timeline: timeline,
	}, nil
}

func NewSense(dir string) (*Sense, error) {
	s, err := sense.NewFromDir(dir)
	if err != nil {
		return nil, err
	}

	return &Sense{senseAPI: s}, nil
}
