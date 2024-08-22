# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a lightweight image for running the app
FROM gcr.io/distroless/base-debian11

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Command to run the executable
CMD ["/app/main"]
