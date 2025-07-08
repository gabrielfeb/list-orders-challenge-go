## ✨ Funcionalidades

* **Criação de Pedidos**: Disponível via gRPC e GraphQL.
* **Listagem de Pedidos**: Disponível via REST, gRPC e GraphQL.
* **Ambiente Dockerizado**: A aplicação, o banco de dados (MySQL) e as migrações são gerenciados pelo Docker Compose.
* **Migrações Automáticas**: O banco de dados é preparado automaticamente ao iniciar o ambiente.

---

## 🛠️ Tecnologias Utilizadas

* **Linguagem**: Go
* **Banco de Dados**: MySQL
* **Containerização**: Docker & Docker Compose
* **APIs**:
    * gRPC
    * REST (com a biblioteca `go-chi`)
    * GraphQL (com a biblioteca `graphql-go`)
* **Ferramentas**: `golang-migrate` para migrações de banco de dados.

---

## ✅ Pré-requisitos

Antes de começar, garanta que você tenha as seguintes ferramentas instaladas na sua máquina:

1.  **Docker**: [Instruções de instalação](https://docs.docker.com/get-docker/)
2.  **Docker Compose**: Geralmente já vem com o Docker Desktop. [Instruções de instalação](https://docs.docker.com/compose/install/)
3.  **grpcurl**: Uma ferramenta de linha de comando para interagir com servidores gRPC. [Instruções de instalação](https://github.com/fullstorydev/grpcurl#installation)

---

## 🚀 Executando o Projeto

Com os pré-requisitos instalados, siga os passos abaixo para colocar toda a aplicação no ar.

### 1. Clone o Repositório

```bash
git clone [https://github.com/gabrielfeb/list-orders-challenge-go.git](https://github.com/gabrielfeb/list-orders-challenge-go.git)
cd list-orders-challenge-go
```

### 2. Inicie o Ambiente

O comando a seguir irá construir a imagem da sua aplicação, iniciar os contêineres do banco de dados e da aplicação, e rodar as migrações para criar a tabela `orders`.

```bash
docker-compose up --build
```

* A flag `--build` é importante na primeira vez para garantir que a imagem Docker seja construída com o código mais recente.
* Aguarde até que os logs se estabilizem e você veja as mensagens indicando que os servidores estão rodando.
* Para rodar em segundo plano no futuro, você pode usar `docker-compose up -d`.

---

## 🔬 Testando as APIs

Com a aplicação rodando, você pode testar todas as interfaces. Lembre-se que você precisa primeiro **criar um pedido** para que as listagens retornem algum dado.

### Portas dos Serviços

| Serviço  | Porta     | Endpoint Principal        |
| :------- | :-------- | :------------------------ |
| REST     | `8080`    | `/order`                  |
| gRPC     | `50051`   | `pb.OrderService`         |
| GraphQL  | `8082`    | `/query`                  |

### 📲 gRPC (Porta: 50051)

Abra um **novo terminal** para executar estes comandos.

#### Criar um Pedido
```bash
grpcurl -plaintext -d '{"price": "199.99", "tax": "10.50"}' localhost:50051 pb.OrderService.CreateOrder
```

#### Listar Pedidos
```bash
grpcurl -plaintext localhost:50051 pb.OrderService.ListOrders
```

### 🌐 REST (Porta: 8080)

#### Listar Pedidos
Você pode usar o cURL, Postman, ou simplesmente abrir o endereço no seu navegador.
```bash
curl http://localhost:8080/order
```

### ⚛️ GraphQL (Porta: 8082)

Acesse a interface gráfica **GraphiQL** no seu navegador, que é ideal para testes:
**[http://localhost:8082/query](http://localhost:8082/query)**

#### Criar um Pedido (Mutation)
Cole o código abaixo no painel da esquerda e execute.
```graphql
mutation {
  createOrder(input: {price: 550.00, tax: 45.75}) {
    id
    price
    tax
    final_price
  }
}
```

#### Listar Pedidos (Query)
Após criar um pedido, execute esta query para visualizá-lo.
```graphql
query {
  listOrders {
    id
    price
    tax
  }
}
```

---

### 🛑 Parando o Ambiente

Para parar todos os contêineres, volte ao terminal onde o `docker-compose` está rodando e pressione `Ctrl + C`. Se estiver rodando em modo detached, use o comando:

```bash
docker-compose down
```