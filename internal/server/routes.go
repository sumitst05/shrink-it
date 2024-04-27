package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./dist"))
	router.HandleFunc("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/dist/", fs).ServeHTTP(w, r)
	})

	return router
}
