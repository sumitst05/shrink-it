package handler

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"os/exec"
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

func HandlePDF() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename=\"shrinked.pdf\"")

		cmd := exec.Command("gs", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dPDFSETTINGS=/ebook", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile=-", "-")
		cmd.Stdin = file
		cmd.Stdout = w
		err = cmd.Run()
		if err != nil {
			http.Error(w, "Could not shrink PDF", http.StatusInternalServerError)
			return
		}
	}
}

func HandleDOCX() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		tmpFile, err := os.CreateTemp("", "original-*.docx")
		if err != nil {
			http.Error(w, "Could not create temporary file", http.StatusInternalServerError)
			return
		}
		defer os.Remove(tmpFile.Name())

		_, err = io.Copy(tmpFile, file)
		if err != nil {
			http.Error(w, "Could not save DOCX file", http.StatusInternalServerError)
			return
		}
		tmpFile.Close()

		cmd := exec.Command("pandoc", "--standalone", tmpFile.Name(), "-o", "shrinked.docx")
		err = cmd.Run()
		if err != nil {
			http.Error(w, "Could not shrink DOCX file", http.StatusInternalServerError)
			return
		}

		shrinkedDocx, err := os.ReadFile("shrinked.docx")
		if err != nil {
			http.Error(w, "Could not read converted DOCX file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.docx"`)
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

		_, err = w.Write(shrinkedDocx)
		if err != nil {
			http.Error(w, "Could not write shrinked DOCX file", http.StatusInternalServerError)
			return
		}
		defer os.Remove("shrinked.docx")
	}
}
