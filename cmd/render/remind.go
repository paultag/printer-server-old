package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Event struct {
	When        time.Time
	Description string
}

type RemindData struct {
	Events []Event
}

type Remind struct {
	CalendarURL      string
	CalendarUsername string
	CalendarPassword string
}

func (Remind) Config() CardConfig {
	return CardConfig{Template: "remind"}
}

func sameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (r Remind) Query() (interface{}, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", r.CalendarURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(r.CalendarUsername, r.CalendarPassword)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(resp.Body)

	reminders := &RemindData{
		Events: []Event{},
	}

	today := time.Now()

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		if line == "" {
			break
		}

		els := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(els) != 2 {
			return nil, fmt.Errorf("Expected two pairs :(")
		}

		when, err := time.Parse("2006/01/02", els[0])
		if err != nil {
			return nil, err
		}

		if !sameDay(today, when) {
			continue
		}

		reminders.Events = append(reminders.Events, Event{
			When:        when,
			Description: els[1],
		})

		if err == io.EOF {
			break
		}
	}

	return &reminders, nil
}

func NewRemind(url, username, password string) (*Remind, error) {
	return &Remind{
		CalendarURL:      url,
		CalendarUsername: username,
		CalendarPassword: password,
	}, nil
}
