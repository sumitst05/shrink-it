package handler

import (
	"net/http"

	"github.com/sumitst05/shrink-it/web/templates/pages"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homepage := pages.Home()
		if err := homepage.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
