# List Orders Challenge Go (v2)

Este projeto implementa um serviço de criação e listagem de ordens com Go, utilizando uma arquitetura avançada e robusta.

## Tecnologias

- **Go**
- **Clean Architecture**
- **Viper:** Gerenciamento de configuração.
- **Google Wire:** Injeção de Dependência automatizada.
- **RabbitMQ:** Comunicação assíncrona via eventos.
- **REST** com `go-chi`.
- **gRPC**
- **GraphQL** com `gqlgen`.
- **PostgreSQL**
- **Docker** & **Docker Compose**

## Como Executar

1.  **Pré-requisitos:** Docker e Docker Compose instalados.
2.  **Clone o repositório.**
3.  **Execute o comando:**

    ```bash
    docker-compose up --build
    ```

Este comando irá iniciar todos os serviços: a aplicação Go, o PostgreSQL e o RabbitMQ.

## Serviços e Portas

- **API REST & GraphQL:** `http://localhost:8080`
- **Servidor gRPC:** `localhost:50051`
- **RabbitMQ Management UI:** `http://localhost:15672` (login: `guest` / `guest`)

## Fluxo de Evento
Ao criar uma nova ordem (via REST ou GraphQL), um evento `OrdemCriada` é publicado no RabbitMQ. Um *consumer* (ainda a ser implementado como exemplo) pode escutar este evento para processamento assíncrono (ex: enviar email, notificar outro microsserviço, etc.).