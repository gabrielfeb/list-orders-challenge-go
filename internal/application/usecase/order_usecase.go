package usecase

import "github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"

type OrderRepository interface {
	Save(order *entity.Order) error
	FindAll() ([]entity.Order, error)
}
type OrderUseCase struct{ OrderRepository OrderRepository }

func NewOrderUseCase(repo OrderRepository) *OrderUseCase { return &OrderUseCase{OrderRepository: repo} }
func (uc *OrderUseCase) CreateOrder(price, tax float64) (*entity.Order, error) {
	order, err := entity.NewOrder(price, tax)
	if err != nil {
		return nil, err
	}
	order.CalculateFinalPrice()
	return order, uc.OrderRepository.Save(order)
}
func (uc *OrderUseCase) ListOrders() ([]entity.Order, error) { return uc.OrderRepository.FindAll() }
