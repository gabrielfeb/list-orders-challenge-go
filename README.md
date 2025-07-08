# List Orders Challenge - Go

Este projeto implementa um serviço em Go para criar e listar pedidos, expondo a funcionalidade através de três interfaces distintas: REST, gRPC e GraphQL.

## ✨ Funcionalidades

* **Criação de Pedidos**: Via gRPC.
* **Listagem de Pedidos**:
    * Endpoint REST: `GET /order`
    * Serviço gRPC: `ListOrders`
    * Query GraphQL: `listOrders`

## 🛠️ Tecnologias

* Go
* Docker & Docker Compose
* MySQL
* gRPC
* GraphQL
* REST (usando Chi)

## 🚀 Como Executar

1.  **Pré-requisitos:**
    * Docker
    * Docker Compose

2.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/gabrielfeb/list-orders-challenge-go.git](https://github.com/gabrielfeb/list-orders-challenge-go.git)
    cd list-orders-challenge-go
    ```

3.  **Suba os contêineres:**
    O comando a seguir irá construir a imagem da aplicação, subir o banco de dados MySQL e rodar as migrações automaticamente.
    ```bash
    docker-compose up --build -d
    ```

4.  **Verifique se os serviços estão rodando:**
    Use `docker-compose ps` para ver o status dos contêineres. Todos devem estar com o status `up` ou `healthy`.

## 🚪 Portas dos Serviços

* **Servidor REST:** `http://localhost:8080`
* **Servidor gRPC:** `localhost:50051`
* **Servidor GraphQL:** `http://localhost:8082`

## 🧪 Testando a API

Você pode usar o arquivo `api.http` com a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) no VS Code ou qualquer outra ferramenta de sua preferência (Postman, Insomnia, cURL).

### Criar um Pedido (via gRPC)
*Para criar um pedido, você precisará de um cliente gRPC como [grpcurl](https://github.com/fullstorydev/grpcurl) ou Postman.*

**Exemplo com grpcurl:**
```bash
grpcurl -plaintext -d '{"price": 250.50, "tax": 25.05}' localhost:50051 pb.OrderService/CreateOrder