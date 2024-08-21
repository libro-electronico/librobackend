# Gunakan base image yang sesuai
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Salin go mod dan sum
COPY go.mod go.sum ./

# Unduh dependensi
RUN go mod download

# Salin source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Gunakan base image yang lebih ringan untuk runtime
FROM gcr.io/distroless/base-debian11

# Salin binary dari builder
COPY --from=builder /app/main /app/main

# Set command untuk menjalankan aplikasi
CMD ["/app/main"]
