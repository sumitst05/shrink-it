package server

import (
	"net/http"

	"github.com/sumitst05/shrink-it/api/handler"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./dist"))
	router.HandleFunc("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/dist/", fs).ServeHTTP(w, r)
	})

	router.HandleFunc("/", handler.HomeHandler())

	router.HandleFunc("/cards", handler.CardHandler())

	return router
}
