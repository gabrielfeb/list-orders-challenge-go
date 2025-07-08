package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielfeb/list-orders-challenge-go/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebOrderHandler struct {
	ListOrdersUseCase usecase.ListOrdersUseCase
}

func NewWebOrderHandler(listUC usecase.ListOrdersUseCase) *WebOrderHandler {
	return &WebOrderHandler{
		ListOrdersUseCase: listUC,
	}
}

func (h *WebOrderHandler) RegisterRoutes(router *chi.Mux) {
	router.Get("/order", h.ListOrders)
}

func (h *WebOrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func NewWebServer(port string) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	return router
}
