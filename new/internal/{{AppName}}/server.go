package {{AppName}}

import (
	"context"
	"fmt"
	"{{ProjectName}}/internal/{{AppName}}/api/v1/registry"
	"{{ProjectName}}/internal/{{AppName}}/config"
	"{{ProjectName}}/internal/{{AppName}}/model"
	myValidator "{{ProjectName}}/internal/{{AppName}}/validator"
	"{{ProjectName}}/pkg/client/database"
	"{{ProjectName}}/pkg/log"
	"{{ProjectName}}/pkg/validator"
	"net/http"
)

type Server struct {
	Config *config.Cfg
	Server *http.Server
	err    error
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	s.initCfg()
	s.initLog()
	s.initDB(stopCh)
	s.initHttpServer()
	s.initRouter()
	s.initValidator()
	return s.err
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
		log.Info(fmt.Sprintf("Shutdown server"))
		_ = s.Server.Shutdown(ctx)
	}()
	log.Info(fmt.Sprintf("Start listening on %s", s.Server.Addr))
	err = s.Server.ListenAndServe()
	return nil
}

func (s *Server) initCfg() {
	if s.err != nil {
		return
	}
	s.Config, s.err = config.TryLoadFromDisk()
}

func (s *Server) initDB(stopCh <-chan struct{}) {
	if s.err != nil {
		return
	}
	var c *database.Client
	log.Info(fmt.Sprintf("db init"))
	c, s.err = database.NewDatabaseClient(s.Config.Mysql, stopCh)
	log.Info(fmt.Sprintf("db init over"))
	model.MainDB = c.DB()
}

func (s *Server) initHttpServer() {
	if s.err != nil {
		return
	}
	s.Server = new(http.Server)
	s.Server.Addr = s.Config.Server.Addr
}

func (s *Server) initLog() {
	if s.err != nil {
		return
	}
	s.err = log.NewLog(s.Config.Log)
}

func (s *Server) initRouter() {
	if s.err != nil {
		return
	}
	s.Server.Handler = registry.Router
}

func (s *Server) initValidator() {
	if s.err != nil {
		return
	}
	s.err = validator.Init(s.Config.Server.Locale, myValidator.RegisterValidation)
}
