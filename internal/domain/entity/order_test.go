package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(100.0, 10.0)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, 100.0, order.Price)
}

func TestOrder_Validate(t *testing.T) {
	_, err := NewOrder(0, 5)
	assert.Error(t, err)
	assert.Equal(t, "price must be greater than zero", err.Error())

	_, err = NewOrder(10, -1)
	assert.Error(t, err)
	assert.Equal(t, "tax cannot be negative", err.Error())
}

func TestOrder_CalculateFinalPrice(t *testing.T) {
	order, err := NewOrder(100.0, 15.0)
	assert.NoError(t, err)
	order.CalculateFinalPrice()
	assert.Equal(t, 115.0, order.FinalPrice)
}
