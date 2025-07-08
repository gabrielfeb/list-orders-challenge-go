package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"list-orders-challenge-go/internal/infra/web/handler"
)

func NewRouter(orderHandler *handler.OrderHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", orderHandler.ListOrders)
	return r
}
