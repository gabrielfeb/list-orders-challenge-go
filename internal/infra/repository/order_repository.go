package repository

import (
	"context"
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/entity"
)

type OrderRepositoryPostgres struct {
	DB *sql.DB
}

func NewOrderRepositoryPostgres(db *sql.DB) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{DB: db}
}

func (r *OrderRepositoryPostgres) Save(order *entity.Order) error {
	_, err := r.DB.Exec(
		"INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)",
		order.ID, order.Price, order.Tax, order.FinalPrice,
	)
	return err
}

func (r *OrderRepositoryPostgres) ListOrders(ctx context.Context) ([]entity.Order, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
