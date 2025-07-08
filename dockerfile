FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /server .
COPY --from=builder /app/migrations ./migrations
EXPOSE 8080 50051
CMD ["./server"]