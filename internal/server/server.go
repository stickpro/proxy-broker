package server

import (
	"asocks-ws/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.HTTPConfig) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
