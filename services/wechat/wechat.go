package wechat

import (
	"log"
	"sync/atomic"

	"github.com/gorilla/mux"
)

type Service struct {
	configValue atomic.Value
	logger      *log.Logger
	router      *mux.Router
}

func NewService(c Config, l *log.Logger, r *mux.Router) *Service {

	s := &Service{
		logger: l,
		router: r,
	}
	b, p, err := s.Login()
	if err != nil {
		s.logger.Printf("E! login fail %s", err.Error())
	}
	s.logger.Printf("I! baseRequest %v", b)
	from, err := s.wxInit(b, p)
	if err != nil {
		s.logger.Printf("E! init fail %s", err.Error())
	}
	groupUsername, err := s.getContact(b, p)
	if err != nil {
		s.logger.Printf("E! getContact fail %s", err.Error())
	}
	if groupUsername == "" {
		c.To = "filehelper"
	} else {
		c.To = groupUsername
	}
	// Doesn't work // FIXME: 2017/3/28
	err = s.notify(b, p, from, c.To)
	if err != nil {
		s.logger.Printf("E! notify fail %s", err.Error())
	}
	c.BaseRequest = b
	c.PassTicket = p
	c.From = from
	s.configValue.Store(c)
	s.router.HandleFunc("/wechat", ServiceHandler(c))

	s.logger.Println("I! service started")

	return s
}
