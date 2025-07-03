//go:build wireinject

package main

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/database"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/graph"
	"github.com/google/wire"
)

func InitializeGraphQLServer(db *sql.DB) (*graph.Resolver, error) {
	wire.Build(
		database.NewOrderRepositoryDb,
		wire.Bind(new(usecase.OrderRepository), new(*database.OrderRepositoryDb)),
		usecase.NewOrderUseCase,
		wire.Struct(new(graph.Resolver), "*"),
	)
	return &graph.Resolver{}, nil
}
