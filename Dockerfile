# Stage 1: Build environment
FROM golang:1.24.2-bullseye AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum (if they exist)
COPY go.mod ./
RUN go mod download || true

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o calculator

# Stage 2: Lightweight image
FROM debian:bullseye-slim

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/calculator .

# Entry point (default command)
ENTRYPOINT ["./calculator"]
