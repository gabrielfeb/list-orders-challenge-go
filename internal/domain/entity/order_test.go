package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1, 100, 10)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, float64(100), order.Price)
}
func TestOrder_Validate(t *testing.T) {
	_, err := NewOrder(0, 10, 5)
	assert.Error(t, err, "id is required")
	_, err = NewOrder(1, 0, 5)
	assert.Error(t, err, "price must be greater than zero")
	_, err = NewOrder(1, 10, -1)
	assert.Error(t, err, "tax cannot be negative")
}
func TestOrder_CalculateFinalPrice(t *testing.T) {
	order, _ := NewOrder(1, 100, 15)
	order.CalculateFinalPrice()
	assert.Equal(t, float64(115), order.FinalPrice)
}
