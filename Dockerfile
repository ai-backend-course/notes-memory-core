# ============================
# BUILDER STAGE
# ============================
FROM golang:1.23-alpine AS builder

# Build with CGO disabled for static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the application
RUN go build -o app .

# ============================
# RUNTIME STAGE
# ============================
FROM alpine:latest

# Install SSL certs (required for HTTPS)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Expose API port
EXPOSE 8080

# Recommended Fiber setting
ENV FIBER_PREFORK=false

# Run application
CMD ["./app"]
