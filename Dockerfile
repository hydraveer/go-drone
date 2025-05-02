# Stage 1: Build the Go binary
FROM golang:1.24.2 AS builder

WORKDIR /app

ENV TMPDIR=/app/tmp
RUN mkdir -p /app/tmp

# Copy only go.mod
COPY go.mod ./
RUN go mod download

# Copy Go source files
COPY *.go ./

# Build the Go app
RUN go build -o calculator

# Stage 2: Lightweight runtime image
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/calculator .

ENTRYPOINT ["./calculator"]
