package event

import (
	"encoding/json"
	"errors"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
	"github.com/streadway/amqp"
)

// ErrHandlerAlreadyRegistered é retornado quando um handler já foi registrado para um evento.
var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// EventDispatcher é a implementação concreta de um despachante de eventos usando RabbitMQ.
type EventDispatcher struct {
	Handlers map[string][]event.EventHandlerInterface
}

// NewEventDispatcher é o construtor para o EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Handlers: make(map[string][]event.EventHandlerInterface),
	}
}

// Dispatch publica um evento em uma exchange do RabbitMQ.
// A implementação atual está simplificada para demonstração.
func (d *EventDispatcher) Dispatch(event event.EventInterface) error {
	// A conexão com RabbitMQ deve ser gerenciada de forma mais robusta em produção (ex: pool de conexões).
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(event.GetPayload())
	if err != nil {
		return err
	}

	// Declara a exchange do tipo 'fanout', que envia para todas as filas ligadas a ela.
	err = ch.ExchangeDeclare(event.GetName(), "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	// Publica a mensagem na exchange.
	err = ch.Publish(event.GetName(), "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}
	return nil
}

// Register registra um handler para um nome de evento.
func (d *EventDispatcher) Register(eventName string, handler event.EventHandlerInterface) error {
	if _, ok := d.Handlers[eventName]; ok {
		for _, h := range d.Handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	d.Handlers[eventName] = append(d.Handlers[eventName], handler)
	return nil
}

// As funções abaixo são parte do contrato, mas não são implementadas neste exemplo.

func (d *EventDispatcher) Remove(eventName string, handler event.EventHandlerInterface) error {
	return nil // Implementação futura
}
func (d *EventDispatcher) Has(eventName string, handler event.EventHandlerInterface) bool {
	return false // Implementação futura
}
func (d *EventDispatcher) Clear() {
	d.Handlers = make(map[string][]event.EventHandlerInterface)
}
