package server

import (
	"context"
	"net/http"
	"time"

	"restplaceholder/internal/pkg/signal"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	go signal.ListenSignals()

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
