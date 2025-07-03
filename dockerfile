# Estágio 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instala as ferramentas essenciais
RUN apk add --no-cache git gcc libc-dev

# Copia os arquivos de módulo e baixa as dependências
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copia todo o código fonte
COPY . .

# Gera o código do Wire
RUN go generate ./cmd/api/...

# Compila o projeto e coloca o executável na raiz do WORKDIR (/app)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# Estágio 2: Final
FROM alpine:latest

WORKDIR /app

# Copia o executável do estágio de build para a imagem final
COPY --from=builder /app/main .

# Copia o arquivo de configuração
COPY configs.yml .

EXPOSE 8080

# Define o comando para rodar a aplicação
CMD ["./main"]