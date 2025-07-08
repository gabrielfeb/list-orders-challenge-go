package usecase

import (
	"context"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: orderRepository}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]ListOrdersOutputDTO, error) {
	orders, err := uc.OrderRepository.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	var output []ListOrdersOutputDTO
	for _, order := range orders {
		output = append(output, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return output, nil
}
