# Use the official Golang image for the build stage
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o librobackend ./run/main.go

# Use a minimal base image to run the Go app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go app from the builder stage
COPY --from=builder /app/librobackend .

# Command to run the Go app
CMD ["./librobackend"]
