# Stage 1: Build the Go binary
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Set custom temp dir to avoid no space error
ENV TMPDIR=/app/tmp
RUN mkdir -p /app/tmp

# Copy go.mod and go.sum
COPY go.mod
RUN go mod download

# Copy only required Go source files
COPY *.go ./

# Build the Go app
RUN go build -o calculator

# Stage 2: Lightweight runtime image
FROM debian:bullseye-slim

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/calculator .

ENTRYPOINT ["./calculator"]
