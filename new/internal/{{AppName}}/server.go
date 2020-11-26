package {{AppName}}

import (
	"context"
	"fmt"
	"net/http"
	"{{ProjectName}}/internal/{{AppName}}/config"
	"{{ProjectName}}/internal/{{AppName}}/model"
	"{{ProjectName}}/pkg/client/database"
	"{{ProjectName}}/pkg/log"
)

type Server struct {
	Config *config.Cfg
	Server *http.Server
	err    error
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	//配置
	s.initCfg()
	s.initDB(stopCh)
	s.initHttpServer()
	s.initLog()
	return s.err
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
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
	c, s.err = database.NewDatabaseClient(s.Config.Mysql, stopCh)
	model.MainDB = c.DB()
	s.migrate()

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

func (s *Server) migrate() {
	model.MainDB.AutoMigrate(
		//new(model.App),
	)
}
