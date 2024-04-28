package main

import (
	"log/slog"
	"os"
	"sync"

	_ "github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/sumitst05/shrink-it/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	server := server.NewServer(port)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := server.Run(); err != nil {
			slog.Error("Failed to start server:", err)
		}
	}()

	go func() {
		defer wg.Done()
		server.SelfPing()
	}()

	wg.Wait()
}
