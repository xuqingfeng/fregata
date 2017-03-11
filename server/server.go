package server

import "fmt"

type Server struct {
	config *Config
}

func New(c *Config) (*Server, error) {

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	s := &Server{
		config: c,
	}

}

func (s *Server) appendSlackService() {

	c := s.config.Slack

}
