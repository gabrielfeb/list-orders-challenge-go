# Orders Challenge

## Serviços
- REST API: 8080
- gRPC: 50051
- GraphQL: 8080/graphql

## Como executar
1. `docker-compose up -d`
2. Acesse os endpoints:
   - REST: http://localhost:8080
   - gRPC: localhost:50051
   - GraphQL: http://localhost:8080/graphql

## Testes
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out