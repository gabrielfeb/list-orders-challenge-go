//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/streadway/amqp"

	"github.com/gabrielfeb/list-orders-challenge-go/configs"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/database"
	infra_event "github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/event"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/web/handler"
)

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

type AllServices struct {
	WebServerHandler *handler.WebOrderHandler
}
