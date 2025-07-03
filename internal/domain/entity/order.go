package entity

import "errors"

type Order struct{ ID, Price, Tax, FinalPrice float64 }

func NewOrder(id float64, price, tax float64) (*Order, error) {
	order := &Order{ID: id, Price: price, Tax: tax}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}
func (o *Order) Validate() error {
	if o.ID == 0 {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax < 0 {
		return errors.New("tax cannot be negative")
	}
	return nil
}
func (o *Order) CalculateFinalPrice() { o.FinalPrice = o.Price + o.Tax }
