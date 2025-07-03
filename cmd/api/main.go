package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabrielfeb/list-orders-challenge-go/configs"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infrastructure/graph"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id SERIAL PRIMARY KEY, price DECIMAL(10, 2) NOT NULL, tax DECIMAL(10, 2) NOT NULL, final_price DECIMAL(10, 2) NOT NULL);")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Inicializa o servidor com o banco de dados
	// e injeta as dependências necessárias
	server, err := InitializeServer(db)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	// Servidor GraphQL usa o resolver de dentro do nosso server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: server.GraphQLResolver}))

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Rotas REST usam o handler de dentro do nosso server
	router.Post("/orders", server.OrderHandler.CreateOrder)
	router.Get("/orders", server.OrderHandler.GetOrders)

	// Rotas GraphQL
	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", srv)

	fmt.Printf("Server is running on port %s\n", cfg.WebServerPort)
	fmt.Printf("GraphQL Playground available at http://localhost:%s/\n", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, router)
}
