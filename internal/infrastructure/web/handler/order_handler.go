package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/dto"
	"github.com/gabrielfeb/list-orders-challenge-go/internal/application/usecase"
	"github.com/google/uuid"
)

// WebOrderHandler é responsável por receber as requisições HTTP para ordens,
// decodificar os dados, chamar os casos de uso apropriados e
// codificar as respostas de volta para o cliente.
type WebOrderHandler struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

// NewWebOrderHandler é o construtor para o WebOrderHandler.
func NewWebOrderHandler(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrdersUseCase *usecase.ListOrdersUseCase,
) *WebOrderHandler {
	return &WebOrderHandler{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

// CreateOrder trata as requisições POST para criar uma nova ordem.
func (h *WebOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dto.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Gera um ID único para a nova ordem.
	input.ID = uuid.New().String()

	output, err := h.CreateOrderUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// GetOrders trata as requisições GET para listar todas as ordens.
func (h *WebOrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
