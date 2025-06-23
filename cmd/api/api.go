//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/streadway/amqp"

	"go-challenge/configs"
	"go-challenge/internal/application/event"
	"go-challenge/internal/application/repository"
	"go-challenge/internal/application/usecase"
	"go-challenge/internal/infrastructure/database"
	infra_event "go-challenge/internal/infrastructure/event"
	graphql_resolver "go-challenge/internal/infrastructure/graph"
	"go-challenge/internal/infrastructure/grpc/service"
	"go-challenge/internal/infrastructure/web/handler"
)

// Providers para eventos
var setEventDispatcherDependency = wire.NewSet(
	infra_event.NewEventDispatcher,
	wire.Bind(new(event.EventDispatcherInterface), new(*infra_event.EventDispatcher)),
)

// Providers para o banco de dados
var setRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(repository.OrderRepository), new(*database.OrderRepository)),
)

// Providers para os Use Cases
var setUseCaseDependency = wire.NewSet(
	usecase.NewCreateOrderUseCase,
	usecase.NewListOrdersUseCase,
)

// Função para construir o WebOrderHandler
func NewWebOrderHandler(db *sql.DB, eventDispatcher event.EventDispatcherInterface, orderCreatedEvent event.EventInterface) *handler.WebOrderHandler {
	// O Wire irá injetar as dependências automaticamente aqui
	createOrderUseCase := usecase.NewCreateOrderUseCase(
		database.NewOrderRepository(db),
		eventDispatcher,
		orderCreatedEvent,
	)
	listOrdersUseCase := usecase.NewListOrdersUseCase(database.NewOrderRepository(db))
	return handler.NewWebOrderHandler(createOrderUseCase, listOrdersUseCase)
}

// Função para construir o gRPC Service
func NewGRPCService(db *sql.DB, eventDispatcher event.EventDispatcherInterface, orderCreatedEvent event.EventInterface) *service.OrderService {
	createOrderUseCase := usecase.NewCreateOrderUseCase(
		database.NewOrderRepository(db),
		eventDispatcher,
		orderCreatedEvent,
	)
	listOrdersUseCase := usecase.NewListOrdersUseCase(database.NewOrderRepository(db))
	return service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
}

// Função para construir o GraphQL Resolver
func NewGraphQLResolver(db *sql.DB, eventDispatcher event.EventDispatcherInterface, orderCreatedEvent event.EventInterface) *graphql_resolver.Resolver {
	createOrderUseCase := usecase.NewCreateOrderUseCase(
		database.NewOrderRepository(db),
		eventDispatcher,
		orderCreatedEvent,
	)
	listOrdersUseCase := usecase.NewListOrdersUseCase(database.NewOrderRepository(db))
	return &graphql_resolver.Resolver{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

// Injetor Principal
func InitializeAllServices(cfg *configs.Config, db *sql.DB, rabbitMQConn *amqp.Connection) (*AllServices, error) {
	wire.Build(
		wire.Struct(new(AllServices), "*"),
		setEventDispatcherDependency,
		NewWebOrderHandler,
		NewGRPCService,
		NewGraphQLResolver,
		// Provider para o evento
		wire.Value(event.EventInterface(&usecase.OrderCreatedEvent{Name: cfg.OrderCreatedEvent})),
	)
	return &AllServices{}, nil
}

type AllServices struct {
	WebServerHandler *handler.WebOrderHandler
	GRPCService      *service.OrderService
	GraphQLResolver  *graphql_resolver.Resolver
}
