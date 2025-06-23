//go:build wireinject
// +build wireinject

//go:generate wire

package main

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/database"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/web/handler"
	"github.com/google/wire"
)

func InitializeOrderHandler(db *sql.DB) *handler.OrderHandler {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(repository.OrderRepository), new(*database.OrderRepository)),
		usecase.NewCreateOrderUseCase,
		usecase.NewListOrdersUseCase,
		handler.NewOrderHandler,
	)
	return &handler.OrderHandler{}
}
