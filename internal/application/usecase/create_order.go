package usecase

import (
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/event"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"
)

type OrderCreatedEvent struct {
	Name    string
	Payload interface{}
}

func (e *OrderCreatedEvent) GetName() string         { return e.Name }
func (e *OrderCreatedEvent) GetPayload() interface{} { return e.Payload }

type CreateOrderUseCase struct {
	OrderRepository   repository.OrderRepository
	EventDispatcher   event.EventDispatcherInterface
	OrderCreatedEvent event.EventInterface
}

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

	eventPayload := &OrderCreatedEvent{
		Name:    c.OrderCreatedEvent.GetName(),
		Payload: order,
	}
	c.EventDispatcher.Dispatch(eventPayload)

	return &dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
