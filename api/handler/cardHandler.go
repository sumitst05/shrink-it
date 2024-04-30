package handler

import (
	"net/http"

	"github.com/sumitst05/shrink-it/web/templates/components"
)

func CardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		choice := r.URL.Query().Get("choice")

		var cards []string
		if choice == "image" {
			cards = append(cards, "png")
			cards = append(cards, "jpg")
		} else if choice == "document" {
			cards = append(cards, "pdf")
			cards = append(cards, "docx")
		} else if choice == "audio" {
			cards = append(cards, "mp3")
			cards = append(cards, "wav")
		} else if choice == "video" {
			cards = append(cards, "mp3")
			cards = append(cards, "wav")
		}

		component := components.Cards(choice, cards)

		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering component", http.StatusInternalServerError)
		}
	}
}
