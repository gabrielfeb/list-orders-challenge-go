package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/configs"
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
	orderHandler := InitializeOrderHandler(db)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/orders", orderHandler.CreateOrder)
	fmt.Printf("Server is running on port %s\n", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, router)
}
