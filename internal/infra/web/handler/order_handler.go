package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase/list_orders"
)

type OrderHandler struct {
	ListOrdersUseCase list_orders.ListOrdersUseCase
}

func NewOrderHandler(listOrdersUseCase list_orders.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
