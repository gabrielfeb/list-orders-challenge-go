package usecase

import (
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"
)

type CreateOrderUseCase struct{ OrderRepository repository.OrderRepository }

func NewCreateOrderUseCase(repo repository.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: repo}
}
func (uc *CreateOrderUseCase) Execute(input dto.OrderInputDTO) (*dto.OrderOutputDTO, error) {
	order, err := entity.NewOrder(0, input.Price, input.Tax) // ID será gerado pelo DB
	if err != nil {
		return nil, err
	}
	order.CalculateFinalPrice()
	if err := uc.OrderRepository.Create(order); err != nil {
		return nil, err
	}
	return &dto.OrderOutputDTO{
		ID: order.ID, Price: order.Price, Tax: order.Tax, FinalPrice: order.FinalPrice,
	}, nil
}
