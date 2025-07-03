package entity

import "errors"

type Order struct{ ID, Price, Tax, FinalPrice float64 }

func NewOrder(price, tax float64) (*Order, error) {
	o := &Order{Price: price, Tax: tax}
	return o, o.Validate()
}
func (o *Order) Validate() error {
	if o.Price <= 0 {
		return errors.New("price must be > 0")
	}
	if o.Tax < 0 {
		return errors.New("tax must be >= 0")
	}
	return nil
}
func (o *Order) CalculateFinalPrice() { o.FinalPrice = o.Price + o.Tax }
