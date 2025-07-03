package entity

import "errors"

type Order struct {
	ID         int
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(price, tax float64) (*Order, error) {
	order := &Order{
		Price: price,
		Tax:   tax,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax < 0 {
		return errors.New("tax cannot be negative")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
