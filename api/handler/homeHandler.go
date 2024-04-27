package handler

import (
	"net/http"

	"github.com/sumitst05/shrink-it/pages"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homepage := pages.Home()
		if err := homepage.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering home page", http.StatusInternalServerError)
			return
		}
	}
}
