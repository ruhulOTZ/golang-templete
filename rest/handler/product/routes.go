package product

import (
	"expenseTracker/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts), middleware.Auth))
	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("GET /products/{id}", http.HandlerFunc(h.GetProductByID))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), middleware.Auth))
}
