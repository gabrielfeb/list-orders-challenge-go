package usecase

import (
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/repository"
)

type ListOrdersUseCase struct {
	OrderRepository repository.OrderRepository
}

func NewListOrdersUseCase(repo repository.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: repo}
}

func (uc *ListOrdersUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	orders, err := uc.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []dto.OrderOutputDTO
	for _, order := range orders {
		output = append(output, dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return output, nil
}
