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

	router.HandleFunc("GET /cards", handler.CardHandler())

	router.HandleFunc("POST /shrink/png", handler.HandlePNG())
	router.HandleFunc("POST /shrink/jpg", handler.HandleJPG())
	router.HandleFunc("POST /shrink/pdf", handler.HandlePDF())
	router.HandleFunc("POST /shrink/docx", handler.HandleDOCX())
	router.HandleFunc("POST /shrink/mp3", handler.HandleMP3())
	router.HandleFunc("POST /shrink/wav", handler.HandleWAV())
	router.HandleFunc("POST /shrink/mp4", handler.HandleMP4())
	router.HandleFunc("POST /shrink/mkv", handler.HandleMKV())

	return router
}
