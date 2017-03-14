package slack

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
	s.configValue.Store(c)
	s.router.HandleFunc("/slack", ServiceHandler(c))

	return s
}

func (s *Service) config() Config {

	return s.configValue.Load().(Config)
}
