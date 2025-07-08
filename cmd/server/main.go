package main

import (
	"log"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/db"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/graphql"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/grpc"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/web"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/web/handlers"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Database
	postgres := db.NewPostgres()
	orderRepo := repository.NewOrderRepositoryPostgres(postgres)

	// Use Cases
	createOrderUC := usecase.NewCreateOrderUseCase(orderRepo)
	listOrdersUC := usecase.NewListOrdersUseCase(orderRepo)

	// Web Server
	webServer := web.NewWebServer(":8080")
	webServer.Router.Use(middleware.Logger)

	// Handlers
	orderHandler := handlers.NewOrderHandler(createOrderUC, listOrdersUC)
	webServer.AddHandler("/order", orderHandler.CreateOrder)
	webServer.AddHandler("/orders", orderHandler.ListOrders)

	// GraphQL
	graphQLHandler := graphql.NewGraphQLHandler(listOrdersUC)
	webServer.Router.Handle("/graphql", graphQLHandler)

	// gRPC
	go grpc.StartGRPCServer("50051", listOrdersUC)

	log.Println("Server starting on port 8080")
	if err := webServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
