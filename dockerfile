# Estágio 1: Build - ESTRUTURA CORRIGIDA E FINAL
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 1. Copia SÓ os arquivos de módulo primeiro.
COPY go.mod go.sum ./

# 2. Sincroniza e baixa as dependências, criando um go.sum perfeito.
RUN go mod tidy

# 3. AGORA sim, copia o resto do código fonte
COPY . .

# 4. Gera o código do Wire
RUN cd cmd/api && go generate

# 5. Compila a aplicação usando o caminho completo do módulo
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main github.com/gabrielfeb/list-orders-challenge-go/cmd/api


# Estágio 2: Final - Ambiente de execução leve
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY configs.yml .
EXPOSE 8080
CMD ["./main"]