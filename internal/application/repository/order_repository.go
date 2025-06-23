package repository

import "github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"

// OrderRepository define o contrato que os adaptadores de persistência de dados
// (camada de infraestrutura) devem seguir. A camada de aplicação depende desta
// interface, e não de uma implementação concreta, invertendo a dependência.
type OrderRepository interface {
	Create(order *entity.Order) error
	FindAll() ([]entity.Order, error)
}
