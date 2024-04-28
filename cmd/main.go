package main

import (
	"log/slog"
	"os"

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

	server.Run()
}
