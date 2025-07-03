// Conteúdo completo do arquivo
package database

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (price, tax, final_price) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Usei float64 para o ID da entidade, mas o BD gera int. Fazendo a conversão.
	var id int
	err = stmt.QueryRow(order.Price, order.Tax, order.FinalPrice).Scan(&id)
	if err != nil {
		return err
	}
	order.ID = float64(id)
	return nil
}

// NOVO MÉTODO ADICIONADO AQUI
func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		var id int
		if err := rows.Scan(&id, &o.Price, &o.Tax, &o.FinalPrice); err != nil {
			return nil, err
		}
		o.ID = float64(id)
		orders = append(orders, o)
	}
	return orders, nil
}
