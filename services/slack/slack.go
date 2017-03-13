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
	}
	s.configValue.Store(c)
	s.router = r
	s.router.HandleFunc("/slack", ServiceHandler)
	return s
}
