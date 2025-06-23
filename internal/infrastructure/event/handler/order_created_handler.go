package handler

import (
	"log"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
)

// OrderCreatedHandler é um exemplo de handler que escuta por eventos de ordem criada.
type OrderCreatedHandler struct{}

// Handle é o método que executa a lógica quando um evento é recebido.
func (h *OrderCreatedHandler) Handle(event event.EventInterface, ch chan<- error) {
	// A lógica a ser executada quando uma ordem for criada
	// Ex: Enviar um email, notificar outro sistema, etc.
	// Aqui, vamos apenas logar a informação para prova de conceito.
	log.Printf("EVENTO RECEBIDO: %s, DADOS: %+v", event.GetName(), event.GetPayload())
}
