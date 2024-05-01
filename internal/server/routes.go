package server

import (
	"net/http"
	"path"

	"github.com/sumitst05/shrink-it/api/handler"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./dist"))
	router.HandleFunc("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		// Determine the MIME type based on the file extension
		switch path.Ext(r.URL.Path) {
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		default:
			// For other file types, use the default Content-Type
			w.Header().Set("Content-Type", "application/octet-stream")
		}
		http.StripPrefix("/dist/", fs).ServeHTTP(w, r)
	})

	router.HandleFunc("/", handler.HomeHandler())

	router.HandleFunc("/cards", handler.CardHandler())

	router.HandleFunc("/shrink/png", handler.HandlePNG())
	router.HandleFunc("/shrink/jpg", handler.HandleJPG())
	router.HandleFunc("/shrink/pdf", handler.HandlePDF())
	router.HandleFunc("/shrink/docx", handler.HandleDOCX())
	router.HandleFunc("/shrink/mp3", handler.HandleMP3())
	router.HandleFunc("/shrink/wav", handler.HandleWAV())
	router.HandleFunc("/shrink/mp4", handler.HandleMP4())
	router.HandleFunc("/shrink/mkv", handler.HandleMKV())

	return router
}
