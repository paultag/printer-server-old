package main

import (
	"pault.ag/go/politico"
)

type PoliticoData struct {
	Stories []politico.Story
}

func (p PoliticoData) Story() politico.Story {
	return p.Stories[0]
}

func (p PoliticoData) Headlines() []string {
	ret := []string{}
	for _, el := range p.Stories[:5] {
		ret = append(ret, string(el.Title))
	}
	return ret
}

type Politico struct{}

func (Politico) Config() CardConfig {
	return CardConfig{Template: "politico"}
}

func (Politico) Query() (interface{}, error) {
	stories, err := politico.News()

	return &PoliticoData{
		Stories: stories,
	}, err
}

func NewPolitico() (*Politico, error) {
	return &Politico{}, nil
}
