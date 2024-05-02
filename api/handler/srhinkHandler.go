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
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

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
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

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
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.pdf"`)

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
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

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

		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.docx"`)

		_, err = w.Write(shrinkedDocx)
		if err != nil {
			http.Error(w, "Could not write shrinked DOCX file", http.StatusInternalServerError)
			return
		}
		defer os.Remove("shrinked.docx")
	}
}

func HandleMP3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "audio/mpeg")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.mp3"`)

		cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-ar", "16000", "-b:a", "32000", "-ac", "1", "-f", "mp3", "-loglevel", "quiet", "pipe:1")
		cmd.Stdin = file
		cmd.Stdout = w

		if err := cmd.Run(); err != nil {
			http.Error(w, "Could not process MP3", http.StatusInternalServerError)
			return
		}
	}
}

func HandleWAV() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "audio/wav")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.mp3"`)

		cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-ar", "16000", "-b:a", "32000", "-ac", "1", "-f", "wav", "-loglevel", "quiet", "pipe:1")
		cmd.Stdin = file
		cmd.Stdout = w

		if err := cmd.Run(); err != nil {
			http.Error(w, "Could not process WAV", http.StatusInternalServerError)
			return
		}
	}
}

func HandleMP4() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		tempFile, err := os.CreateTemp("", "original-*.mp4")
		if err != nil {
			http.Error(w, "Could not create temporary file", http.StatusInternalServerError)
			return
		}
		defer os.Remove(tempFile.Name())

		_, err = io.Copy(tempFile, file)
		if err != nil {
			http.Error(w, "Could not save uploaded file", http.StatusInternalServerError)
			return
		}
		tempFile.Close()

		outputFile := "shrinked.mp4"
		cmd := exec.Command("ffmpeg", "-i", tempFile.Name(), "-c:v", "libx265", "-preset", "veryfast", "-tag:v", "hvc1", "-b:v", "800k", "-bufsize", "1200k", "-vf", "scale=1080:1920,format=yuv420p", "-b:a", "128k", outputFile)

		if err := cmd.Run(); err != nil {
			http.Error(w, "Could not process MP4", http.StatusInternalServerError)
			return
		}

		processedFile, err := os.Open(outputFile)
		if err != nil {
			http.Error(w, "Could not open processed file", http.StatusInternalServerError)
			return
		}
		defer processedFile.Close()

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.mp4"`)

		_, err = io.Copy(w, processedFile)
		if err != nil {
			http.Error(w, "Could not write processed file", http.StatusInternalServerError)
			return
		}
		defer os.Remove("shrinked.mp4")
	}
}

func HandleMKV() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "video/x-matroska")
		w.Header().Set("Content-Disposition", `attachment; filename="shrinked.mkv"`)

		cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx265", "-preset", "veryfast", "-tag:v", "hvc1", "-b:v", "800k", "-bufsize", "1200k", "-vf", "scale=1080:1920,format=yuv420p", "-b:a", "128k", "-f", "matroska", "-loglevel", "quiet", "pipe:1")
		cmd.Stdin = file
		cmd.Stdout = w

		if err := cmd.Run(); err != nil {
			http.Error(w, "Could not process MKV", http.StatusInternalServerError)
			return
		}
	}
}
