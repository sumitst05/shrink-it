package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	server := http.Server{
		Addr:           s.addr,
		Handler:        s.RegisterRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Info("Server listening on port", "port", os.Getenv("PORT"))

	return server.ListenAndServe()
}
