package usecase

import (
	"github.com/gabrielfeb/list-orders-challenge-go/internal/entity"
)

type CreateOrderInputDTO struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: orderRepository}
}

func (uc *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (*CreateOrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = uc.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutputDTO{
		ID:       