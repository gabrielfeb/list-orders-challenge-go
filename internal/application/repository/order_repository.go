package repository

import "github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"

type OrderRepository interface {
	Create(order *entity.Order) error
}
