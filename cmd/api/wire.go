//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/configs"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/database"
	infra_event "github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/event"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/web/handler"
	"github.com/streadway/amqp"
	"github.comcom/google/wire"
)

// Este arquivo define as injeções de dependência para o Google Wire.
// O Wire irá ler este arquivo e gerar o código de inicialização em 'wire_gen.go'.

var setRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(repository.OrderRepository), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	infra_event.NewEventDispatcher,
	wire.Bind(new(event.EventDispatcherInterface), new(*infra_event.EventDispatcher)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher event.EventDispatcherInterface, event event.EventInterface) *usecase.CreateOrderUseCase {
	return usecase.NewCreateOrderUseCase(database.NewOrderRepository(db), eventDispatcher, event)
}

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	return usecase.NewListOrdersUseCase(database.NewOrderRepository(db))
}

func NewWebOrderHandler(createUsecase *usecase.CreateOrderUseCase, listUsecase *usecase.ListOrdersUseCase) *handler.WebOrderHandler {
	return handler.NewWebOrderHandler(createUsecase, listUsecase)
}

// AllServices é uma struct para agrupar todos os serviços inicializados.
type AllServices struct {
	WebServerHandler *handler.WebOrderHandler
}

// InitializeAllServices é o injetor principal que o Wire usará para construir o grafo de dependências.
func InitializeAllServices(cfg *configs.Config, db *sql.DB, rabbitMQConn *amqp.Connection) (*AllServices, error) {
	wire.Build(
		setEventDispatcherDependency,
		NewCreateOrderUseCase,
		NewListOrderUseCase,
		NewWebOrderHandler,
		wire.Bind(new(event.EventInterface), new(*usecase.OrderCreatedEvent)),
		wire.Value(&usecase.OrderCreatedEvent{Name: cfg.OrderCreatedEvent}),
		wire.Struct(new(AllServices), "*"),
	)
	return &AllServices{}, nil
}
