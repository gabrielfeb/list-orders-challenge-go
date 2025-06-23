//go:build wireinject

package main

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/database"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/web/handler"
	"github.com/google/wire"
)

func NewCreateOrderUseCase(db *sql.DB) *usecase.CreateOrderUseCase {
	repo := database.NewOrderRepository(db)
	return usecase.NewCreateOrderUseCase(repo)
}
func InitializeOrderHandler(db *sql.DB) *handler.OrderHandler {
	wire.Build(
		database.NewOrderRepository,
		wire.Bind(new(repository.OrderRepository), new(*database.OrderRepository)),
		usecase.NewCreateOrderUseCase,
		handler.NewOrderHandler,
	)
	return &handler.OrderHandler{}
}
