package usecase

import "github.com/gabrielfeb/list-orders-challenge-go/internal/entity"

type ListOrdersOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]*ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var output []*ListOrdersOutputDTO
	for _, order := range orders {
		output = append(output, &ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return output, nil
}
