# Estágio 1: Build - Ambiente de compilação com todas as ferramentas
FROM golang:1.24-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Instala as dependências do sistema operacional necessárias para compilar
# git: para baixar dependências Go
# gcc/libc-dev: para compilação CGO (se necessário)
RUN apk add --no-cache git gcc libc-dev protobuf-dev

# Instala as ferramentas Go para gRPC (se você for readicionar no futuro)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

# Copia os arquivos de gerenciamento de dependências
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia todo o código fonte do projeto para o container
COPY . .

# Gera o código do Wire (essencial!)
RUN cd cmd/api && go generate

# Compila a aplicação Go
# -o main: define o nome do arquivo de saída como 'main'
# CGO_ENABLED=0: desabilita CGO para criar um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main github.com/gabrielfeb/list-orders-challenge-go/cmd/api

# Estágio 2: Final - Ambiente de execução leve
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /app

# Copia APENAS o binário compilado do estágio 'builder'
COPY --from=builder /app/main .

# Copia o arquivo de configuração para que a aplicação possa lê-lo
COPY configs.yml .

# Expõe a porta que o nosso servidor web usa
EXPOSE 8080

# Comando que será executado quando o container iniciar
CMD ["./main"]