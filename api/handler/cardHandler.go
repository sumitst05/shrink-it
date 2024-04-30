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
			cards = append(cards, "PNG")
			cards = append(cards, "JPG")
		} else if choice == "document" {
			cards = append(cards, "PDF")
			cards = append(cards, "DOCX")
		} else if choice == "audio" {
			cards = append(cards, "MP3")
			cards = append(cards, "WAV")
		} else if choice == "video" {
			cards = append(cards, "MP4")
			cards = append(cards, "MKV")
		}

		component := components.Cards(choice, cards)

		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering component", http.StatusInternalServerError)
		}
	}
}
