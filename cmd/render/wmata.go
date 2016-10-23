package main

import (
	"pault.ag/go/wmata"
)

type WMATAData struct {
	Incidents []wmata.Incident
}

func (data WMATAData) GreenLineIncidents() []wmata.Incident {
	line := wmata.GreenLine
	ret := []wmata.Incident{}

	for _, el := range data.Incidents {
		for _, affectedLine := range el.LinesAffected {
			if affectedLine == line {
				ret = append(ret, el)
			}
		}
	}

	return ret
}

type WMATA struct {
	Lines []wmata.Line
}

func (WMATA) Config() CardConfig {
	return CardConfig{Template: "wmata"}
}

func (WMATA) Query() (interface{}, error) {
	incidents, err := wmata.GetIncidents()
	return &WMATAData{
		Incidents: incidents.Incidents,
	}, err
}

func NewWMATA(apikey string, lines []wmata.Line) (*WMATA, error) {
	wmata.SetAPIKey(apikey)
	return &WMATA{
		Lines: lines,
	}, nil
}
