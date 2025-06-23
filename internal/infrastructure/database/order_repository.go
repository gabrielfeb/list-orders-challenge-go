package database

import (
	"database/sql"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/domain/entity"
)

type OrderRepository struct{ Db *sql.DB }

func NewOrderRepository(db *sql.DB) *OrderRepository { return &OrderRepository{Db: db} }
func (r *OrderRepository) Create(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (price, tax, final_price) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.QueryRow(order.Price, order.Tax, order.FinalPrice).Scan(&order.ID)
}
