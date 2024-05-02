package main

import (
	"os"

	_ "github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/sumitst05/shrink-it/internal/server"
)

func main() {
	var port string
	godotenv.Load()
	port = os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	server := server.NewServer(port)

	server.Run()
}
