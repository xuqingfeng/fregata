package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/services/macos"
	"github.com/xuqingfeng/fregata/services/slack"
	"github.com/xuqingfeng/fregata/vars"
)

type Server struct {
	config *Config

	router *mux.Router

	LogService logging.Interface
	Logger     *log.Logger
}

func New(c *Config, logService logging.Interface) (*Server, error) {

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	l := logService.NewLogger("[server] ", log.LstdFlags)
	router := mux.NewRouter()
	apiRouter := router.PathPrefix(vars.APIBasePath).Subrouter()
	s := &Server{
		config:     c,
		router:     apiRouter,
		LogService: logService,
		Logger:     l,
	}
	s.Logger.Printf("I! %s server started\n", vars.DaemonName)

	s.appendSlackService()
	s.appendMacosService()

	if err := http.ListenAndServe(":2017", router); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return s, nil
}

func (s *Server) appendSlackService() {

	c := s.config.Slack
	if c.Enabled {
		l := s.LogService.NewLogger("[slack] ", log.LstdFlags)
		r := s.router
		slack.NewService(c, l, r)
	}
}

func (s *Server) appendMacosService() {

	c := s.config.Macos
	if c.Enabled {
		l := s.LogService.NewLogger("[macos] ", log.LstdFlags)
		r := s.router
		macos.NewService(c, l, r)
	}
}
