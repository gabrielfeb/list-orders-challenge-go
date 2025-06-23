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
	"github.com/streadway/amqp"
)

// main é a função de entrada da aplicação.
// O fluxo de inicialização é:
// 1. Carregar as configurações com o Viper.
// 2. Estabelecer conexões com serviços externos (Banco de Dados, RabbitMQ).
// 3. Utilizar o Google Wire (através de InitializeAllServices) para construir o grafo de dependências.
// 4. Iniciar os servidores e listeners.
func main() {
	// 1. Carrega as configurações de 'configs.yml' e variáveis de ambiente.
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// 2. Estabelece conexão com o Banco de Dados.
	db, err := sql.Open(cfg.DBDriver, cfg.DBURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// 2. Estabelece conexão com o RabbitMQ.
	rabbitMQ, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// 3. Inicializa todos os serviços e suas dependências usando o injetor gerado pelo Wire.
	services, err := InitializeAllServices(cfg, db, rabbitMQ)
	if err != nil {
		log.Fatalf("Erro ao inicializar serviços com Wire: %v", err)
	}

	// 4. Configura o roteador web e inicia o servidor HTTP.
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/order", services.WebServerHandler.CreateOrder)
	router.Get("/orders", services.WebServerHandler.GetOrders)

	fmt.Printf("Servidor Web rodando na porta %s\n", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, router)
}
