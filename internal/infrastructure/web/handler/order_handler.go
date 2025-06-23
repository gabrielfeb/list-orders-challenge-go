package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
)

type OrderHandler struct{ CreateOrderUseCase *usecase.CreateOrderUseCase }

func NewOrderHandler(uc *usecase.CreateOrderUseCase) *OrderHandler {
	return &OrderHandler{CreateOrderUseCase: uc}
}
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dto.OrderInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.CreateOrderUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
