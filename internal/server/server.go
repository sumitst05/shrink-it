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

	slog.Info("Server listening on", "port", os.Getenv("PORT"))

	return server.ListenAndServe()
}

func (s *Server) SelfPing() {
	ticker := time.NewTicker(14 * time.Minute)
	go func() {
		for range ticker.C {
			s.pingServer()
		}
	}()
}

func (s *Server) pingServer() {
	res, err := http.Get("http://localhost" + s.addr)
	if err != nil {
		slog.Error("Failed to ping server:", err)
		return
	}
	defer res.Body.Close()
	slog.Info("Server pinged successfully")
}
