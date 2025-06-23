package event

// EventDispatcherInterface defines the contract for an event dispatcher
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	Handlers map[string][]event.EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Handlers: make(map[string][]event.EventHandlerInterface),
	}
}

func (d *EventDispatcher) Dispatch(event event.EventInterface) error {
	if handlers, ok := d.Handlers[event.GetName()]; ok {
		// Publicar no RabbitMQ aqui
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/") // Idealmente, a URL vem da config
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

		err = ch.ExchangeDeclare(event.GetName(), "fanout", true, false, false, false, nil)
		if err != nil {
			return err
		}

		err = ch.Publish(event.GetName(), "", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
		if err != nil {
			return err
		}

		// A lógica de executar handlers pode ser adaptada para consumers
		go func() {
			errorChannel := make(chan error)
			for _, handler := range handlers {
				handler.Handle(event, errorChannel)
			}
		}()
	}
	return nil
}

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

func (d *EventDispatcher) Remove(eventName string, handler event.EventHandlerInterface) error {
	// Implementação de remoção
	return nil
}
func (d *EventDispatcher) Has(eventName string, handler event.EventHandlerInterface) bool {
	// Implementação de verificação
	return false
}
func (d *EventDispatcher) Clear() {
	d.Handlers = make(map[string][]event.EventHandlerInterface)
}