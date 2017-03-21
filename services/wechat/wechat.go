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
	c.BaseRequest = b
	c.PassTicket = p
	s.configValue.Store(c)
	s.router.HandleFunc("/wechat", ServiceHandler(c))

	s.logger.Println("I! service started")

	return s
}
