package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/services/macos"
	"github.com/xuqingfeng/fregata/services/slack"
	"github.com/xuqingfeng/fregata/services/smtp"
	"github.com/xuqingfeng/fregata/services/telegram"
	"github.com/xuqingfeng/fregata/services/twilio"
	"github.com/xuqingfeng/fregata/services/wechat"
	"github.com/xuqingfeng/fregata/vars"
)

type Server struct {
	config *Config

	router *mux.Router

	LogService logging.Interface
	Logger     *log.Logger
}

// New start http server and activate relative services
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
	s.Logger.Printf("I! %s started\n", vars.Name)
	s.router.HandleFunc("/ping", ServiceHandler)

	// start services in parallel
	go func() {
		s.appendSlackService()
	}()
	go func() {
		s.appendMacosService()
	}()
	go func() {
		s.appendTelegramService()
	}()
	go func() {
		s.appendTwilioService()
	}()
	go func() {
		s.appendSMTPService()
	}()
	go func() {
		s.appendWechatService()
	}()

	if err := http.ListenAndServe(c.HTTP.BindAddress, router); err != nil {
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

func (s *Server) appendTelegramService() {

	c := s.config.Telegram
	if c.Enabled {
		l := s.LogService.NewLogger("[telegram] ", log.LstdFlags)
		r := s.router
		telegram.NewService(c, l, r)
	}
}

func (s *Server) appendTwilioService() {

	c := s.config.Twilio
	if c.Enabled {
		l := s.LogService.NewLogger("[twilio] ", log.LstdFlags)
		r := s.router
		twilio.NewService(c, l, r)
	}
}

func (s *Server) appendWechatService() {

	c := s.config.Wechat
	if c.Enabled {
		l := s.LogService.NewLogger("[wechat] ", log.LstdFlags)
		r := s.router

		wechat.NewService(c, l, r)
	}
}

func (s *Server) appendSMTPService() {

	c := s.config.SMTP
	if c.Enabled {
		l := s.LogService.NewLogger("[smtp] ", log.LstdFlags)
		r := s.router

		smtp.NewService(c, l, r)
	}
}
