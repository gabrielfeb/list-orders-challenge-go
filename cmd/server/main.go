package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/infra/graph"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	graphQLServer, err := InitializeGraphQLServer(db)
	if err != nil {
		panic(err)
	}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graphQLServer})))
	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
