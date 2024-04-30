package main

import (
	"os"

	_ "github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/sumitst05/shrink-it/internal/server"
)

func main() {
	var port string
	if err := godotenv.Load(); err != nil {
		port = ":3000"
	} else {
		port = os.Getenv("PORT")
	}

	server := server.NewServer(port)

	server.Run()
}
