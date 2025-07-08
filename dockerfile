FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd/server

FROM alpine:latest

WORKDIR /

COPY --from=builder /main .

EXPOSE 8080 50051 8082

CMD ["./main"]