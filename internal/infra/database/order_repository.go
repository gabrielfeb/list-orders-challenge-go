package database

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"
)

type OrderRepositoryDb struct{ Db *sql.DB }

func NewOrderRepositoryDb(db *sql.DB) *OrderRepositoryDb { return &OrderRepositoryDb{Db: db} }
func (r *OrderRepositoryDb) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (price, tax, final_price) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.QueryRow(order.Price, order.Tax, order.FinalPrice).Scan(&order.ID)
}
func (r *OrderRepositoryDb) FindAll() ([]entity.Order, error) {
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
