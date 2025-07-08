package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"list-orders-challenge-go/internal/infra/database"
	"list-orders-challenge-go/internal/infra/graphql"
	service "list-orders-challenge-go/internal/infra/grpc"
	"list-orders-challenge-go/internal/infra/grpc/pb"
	webserver "list-orders-challenge-go/internal/infra/web"
	"list-orders-challenge-go/internal/usecase"
)

func main() {
	// --- Configuração do Banco de Dados ---
	db, err := sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// --- Inicialização do Repositório e Use Cases ---
	orderRepository := database.NewOrderRepository(db)
	createOrderUseCase := *usecase.NewCreateOrderUseCase(orderRepository)
	listOrdersUseCase := *usecase.NewListOrdersUseCase(orderRepository)

	// --- Canais para sincronização e erros ---
	errChan := make(chan error, 3)

	// --- Iniciar Servidor Web (REST) ---
	go startWebServer(os.Getenv("WEB_SERVER_PORT"), listOrdersUseCase, errChan)

	// --- Iniciar Servidor gRPC ---
	go startGRPCServer(os.Getenv("GRPC_SERVER_PORT"), createOrderUseCase, listOrdersUseCase, errChan)

	// --- Iniciar Servidor GraphQL ---
	go startGraphQLServer(os.Getenv("GRAPHQL_SERVER_PORT"), createOrderUseCase, listOrdersUseCase, errChan)

	log.Printf("Application started. Waiting for connections...")
	err = <-errChan
	log.Fatalf("Error running server: %v", err)
}

func startWebServer(port string, listUC usecase.ListOrdersUseCase, errChan chan error) {
	router := webserver.NewWebServer(":" + port)
	webOrderHandler := webserver.NewWebOrderHandler(listUC)
	webOrderHandler.RegisterRoutes(router)

	log.Printf("REST server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		errChan <- fmt.Errorf("web server error: %w", err)
	}
}

func startGRPCServer(port string, createUC usecase.CreateOrderUseCase, listUC usecase.ListOrdersUseCase, errChan chan error) {
	orderService := service.NewOrderService(createUC, listUC)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errChan <- fmt.Errorf("grpc failed to listen: %w", err)
		return
	}

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		errChan <- fmt.Errorf("grpc server error: %w", err)
	}
}

func startGraphQLServer(port string, createUC usecase.CreateOrderUseCase, listUC usecase.ListOrdersUseCase, errChan chan error) {

	resolver := graphql.NewResolver(createUC, listUC)

	http.Handle("/query", resolver.Handler())

	log.Printf("GraphQL server is running on port %s (endpoint /query)", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		errChan <- fmt.Errorf("graphql server error: %w", err)
	}
}
