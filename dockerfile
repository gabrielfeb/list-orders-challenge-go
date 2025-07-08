# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy migrations (if needed, though migrate container is preferred)
# COPY internal/infra/database/sql/migrations ./internal/infra/database/sql/migrations

EXPOSE 8080 50051 8082

CMD ["./main"]