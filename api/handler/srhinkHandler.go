package handler

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

func HandlePNG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
			http.Error(w, "Could not decode image", http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer

		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(&buf, img)
		if err != nil {
			http.Error(w, "Could not encode image", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.png"`)

		_, err = io.Copy(w, &buf)
		if err != nil {
			http.Error(w, "Could not write image", http.StatusInternalServerError)
			return
		}
	}
}

func HandleJPG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		img, err := jpeg.Decode(file)
		if err != nil {
			http.Error(w, "Could not decode image", http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer

		var opt jpeg.Options
		opt.Quality = 80

		err = jpeg.Encode(&buf, img, &opt)
		if err != nil {
			http.Error(w, "Could not encode image", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", `attachment; filename=shrinked.jpg`)

		_, err = io.Copy(w, &buf)
		if err != nil {
			http.Error(w, "Could not write image", http.StatusInternalServerError)
			return
		}
	}
}
