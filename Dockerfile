# Use the official Golang image as the base image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install clang and DKLS libraries for CGO worker build
RUN apt-get update && apt-get install -y clang && rm -rf /var/lib/apt/lists/*

# Build the HTTP server (no CGO needed)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/vultisigner/main.go

# Build the asynq worker (requires CGO for DKLS/Schnorr C bindings)
ENV LD_LIBRARY_PATH=/go/pkg/mod/github.com/vultisig/go-wrappers@v0.0.0-20260223034715-9a5927a3c4c6/includes/linux
RUN CGO_ENABLED=1 CC=clang go build -o worker cmd/worker/main.go

# Start a new stage from scratch — needs glibc 2.39+ for worker CGO binary
FROM debian:trixie-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/worker .
COPY entrypoint.sh .

# Install CA certificates for TLS connections to external relay
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy DKLS/Schnorr shared libraries for the worker
COPY --from=builder /go/pkg/mod/github.com/vultisig/go-wrappers@v0.0.0-20260223034715-9a5927a3c4c6/includes/linux/*.so /usr/local/lib/
RUN ldconfig
RUN chmod +x entrypoint.sh

# Expose port 8080 to the outside world
EXPOSE 8080

ENV ENV=production

# Run both the HTTP server and the asynq worker
CMD ["./entrypoint.sh"]
