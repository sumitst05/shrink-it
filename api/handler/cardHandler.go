package handler

import (
	"net/http"

	"github.com/sumitst05/shrink-it/web/templates/components"
)

func CardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		choice := r.URL.Query().Get("choice")

		component := components.Cards(choice)

		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering component", http.StatusInternalServerError)
		}
	}
}
