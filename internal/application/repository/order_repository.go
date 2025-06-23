package repository

import "github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"

// OrderRepository é o contrato que o nosso banco de dados deve seguir.
type OrderRepository interface {
	Create(order *entity.Order) error
	FindAll() ([]entity.Order, error)
}
