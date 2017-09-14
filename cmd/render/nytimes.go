package main

import (
	"pault.ag/go/nytimes"
)

type NYTimesData struct {
	Stories []nytimes.Article
}

func (p NYTimesData) Story() nytimes.Article {
	return p.Stories[0]
}

func (p NYTimesData) Headlines() []string {
	ret := []string{}
	for _, el := range p.Stories[:5] {
		ret = append(ret, string(el.Title))
	}
	return ret
}

type NYTimes struct {
	APIKey string
}

func (NYTimes) Config() CardConfig {
	return CardConfig{Template: "nytimes"}
}

func (n NYTimes) Query() (interface{}, error) {
	stories, err := nytimes.TopNews(n.APIKey)
	if err != nil {
		return nil, err
	}

	return &NYTimesData{
		Stories: stories.Results,
	}, err
}

func NewNYTimes(apiKey string) (*NYTimes, error) {
	return &NYTimes{APIKey: apiKey}, nil
}
