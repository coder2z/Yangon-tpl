package app

import (
	"{{.ProjectName}}/cmd"
	s "{{.ProjectName}}/internal/{{.AppName}}"
)

func Run(stopCh <-chan struct{}) error {
	server := NewServer()
	err := server.PrepareRun(stopCh)
	if err != nil {
		return err
	}
	return server.Run(stopCh)
}

func NewServer() cmd.App {
	return &s.Server{}
}
