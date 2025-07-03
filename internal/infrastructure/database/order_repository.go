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
	// O Scan agora é diretamente para o order.ID que é um int
	return stmt.QueryRow(order.Price, order.Tax, order.FinalPrice).Scan(&order.ID)
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(&o.ID, &o.Price, &o.Tax, &o.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
