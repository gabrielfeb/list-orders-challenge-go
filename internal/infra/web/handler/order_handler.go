package handler

import (
	"encoding/json"
	"net/http"

	"list-orders-challenge-go/internal/usecase"
)

type OrderHandler struct {
	ListOrdersUseCase usecase.ListOrdersUseCase
}

func NewOrderHandler(listOrdersUseCase usecase.ListOrdersUseCase) *OrderHandler {
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
