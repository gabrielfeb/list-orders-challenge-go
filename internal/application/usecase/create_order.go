package usecase

import (
	"[github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto](https://github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto)"
	"[github.com/gabrielfeb/list-orders-challenge-go/internal/application/event](https://github.com/gabrielfeb/list-orders-challenge-go/internal/application/event)"
	"[github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository](https://github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository)"
	"[github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity](https://github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity)"
)

// OrderCreatedEvent define a estrutura do evento que é disparado quando uma ordem é criada.
type OrderCreatedEvent struct {
	Name    string
	Payload interface{}
}

// GetName retorna o nome do evento.
func (e *OrderCreatedEvent) GetName() string { return e.Name }

// GetPayload retorna os dados associados ao evento.
func (e *OrderCreatedEvent) GetPayload() interface{} { return e.Payload }

// CreateOrderUseCase é o caso de uso responsável por orquestrar a criação de uma nova ordem.
// Ele depende de abstrações (interfaces) para interagir com o mundo externo,
// como o repositório para persistência e o dispatcher para eventos.
type CreateOrderUseCase struct {
	OrderRepository   repository.OrderRepository
	EventDispatcher   event.EventDispatcherInterface
	OrderCreatedEvent event.EventInterface
}

// NewCreateOrderUseCase é o construtor para o CreateOrderUseCase.
func NewCreateOrderUseCase(
	orderRepository repository.OrderRepository,
	eventDispatcher event.EventDispatcherInterface,
	orderCreatedEvent event.EventInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:   orderRepository,
		EventDispatcher:   eventDispatcher,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

// Execute é o método principal que executa a lógica do caso de uso.
// Recebe um DTO de entrada, cria a entidade, persiste no banco através do repositório
// e dispara um evento através do dispatcher.
func (c *CreateOrderUseCase) Execute(input dto.OrderInputDTO) (*dto.OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := order.Validate(); err != nil {
		return nil, err
	}
	if err := c.OrderRepository.Create(&order); err != nil {
		return nil, err
	}

	// Cria e dispara o evento de ordem criada.
	eventPayload := &OrderCreatedEvent{
		Name:    c.OrderCreatedEvent.GetName(),
		Payload: order,
	}
	c.EventDispatcher.Dispatch(eventPayload)

	// Retorna um DTO de saída com os dados da ordem criada.
	return &dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
