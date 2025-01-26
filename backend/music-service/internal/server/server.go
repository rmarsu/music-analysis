package server

import (
	"context"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + os.Getenv("MUSIC_SERVER"),
			Handler: handler,
			
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
