# Start from the official Golang image for building
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the server binary
RUN go build -o server ./cmd/server

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/server .

# Expose the server port (change if your server uses a different port)
EXPOSE 3000

# Run the server
CMD ["./server"]