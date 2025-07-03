package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(price, tax float64) (*Order, error) {
	if price <= 0 {
		return nil, ErrInvalidPrice
	}
	if tax <= 0 {
		return nil, ErrInvalidTax
	}

	return &Order{
		ID:         uuid.New().String(),
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}, nil
}

var (
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidTax   = errors.New("invalid tax")
)
