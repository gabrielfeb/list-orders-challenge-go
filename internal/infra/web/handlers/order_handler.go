package handlers

import (
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	CreateOrderUC *usecase.CreateOrderUseCase
	ListOrdersUC  *usecase.ListOrdersUseCase
}

func NewOrderHandler(createUC *usecase.CreateOrderUseCase, listUC *usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		CreateOrderUC: createUC,
		ListOrdersUC:  listUC,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input struct {
		Price float64 `json:"price"`
		Tax   float64 `json:"tax"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implementação da criação...
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.ListOrdersUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
