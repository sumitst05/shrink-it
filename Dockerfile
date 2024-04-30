FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

# Download the standalone TailwindCSS CLI
RUN curl -L https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o /usr/local/bin/tailwindcss

# Make it executable
RUN chmod +x /usr/local/bin/tailwindcss

# Compile TailwindCSS styles
RUN tailwindcss -i ./web/static/index.css -o ./dist/style.css

# Install Ghostscript, Pandoc, and FFmpeg
RUN apt-get update && \
    apt-get install -y --no-install-recommends ghostscript pandoc ffmpeg && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Build Go app
RUN go build -o main ./cmd/main.go

EXPOSE 3000

CMD ["./main"]
