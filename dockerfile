# Gunakan base image untuk build Go
FROM golang:1.19-alpine AS builder

# Set working directory
WORKDIR /app

# Salin go.mod dan go.sum jika ada
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh kode sumber
COPY . .

# Build aplikasi
RUN go build -o librobackend

# Gunakan base image yang lebih ringan untuk runtime
FROM alpine:latest

# Salin binary dari builder
COPY --from=builder /app/librobackend /app/librobackend

# Set command untuk menjalankan aplikasi
CMD ["/app/librobackend"]
