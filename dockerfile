FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git gcc libc-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go run github.com/99designs/gqlgen generate
RUN cd cmd/server && go generate
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY configs.yml .
COPY internal/infra/database/migration ./internal/infra/database/migration
EXPOSE 8080
CMD ["/app/main"]