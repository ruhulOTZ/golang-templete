package cmd

import (
	"expenseTracker/handler"
	"expenseTracker/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProducts), middleware.Auth))

	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handler.GetProductByID))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))

}
