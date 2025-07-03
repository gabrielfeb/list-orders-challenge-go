//go:build wireinject
// +build wireinject

//go:generate wire
package main

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/database"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/graph"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/web/handler"
	"github.com/google/wire"
)

type Server struct {
	OrderHandler    *handler.OrderHandler
	GraphQLResolver *graph.Resolver
}

func NewServer(orderHandler *handler.OrderHandler, graphqlResolver *graph.Resolver) *Server {
	return &Server{
		OrderHandler:    orderHandler,
		GraphQLResolver: graphqlResolver,
	}
}

func NewGraphQLResolver(createUC *usecase.CreateOrderUseCase, listUC *usecase.ListOrdersUseCase) *graph.Resolver {
	return &graph.Resolver{
		CreateOrderUseCase: createUC,
		ListOrdersUseCase:  listUC,
	}
}

func InitializeServer(db *sql.DB) (*Server, error) {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(repository.OrderRepository), new(*database.OrderRepository)),
		usecase.NewCreateOrderUseCase,
		usecase.NewListOrdersUseCase,
		handler.NewOrderHandler,
		NewGraphQLResolver,
		NewServer,
	)
	return &Server{}, nil
}
