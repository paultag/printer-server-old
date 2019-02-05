package main

import (
	"crypto/tls"

	"pault.ag/go/sidequest/client"
)

type Sidequest struct {
	Client client.API
}

func (s Sidequest) Config() CardConfig {
	return CardConfig{Template: "sidequest"}
}

func (s Sidequest) Query() (interface{}, error) {
	tasks, err := s.Client.GetTasks(map[string][]string{
		"state": []string{"nearlydue"},
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func NewSidequest(apiBase, certPath, keyPath string) (*Sidequest, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	client := client.New(apiBase, cert)
	return &Sidequest{Client: client}, nil
}
