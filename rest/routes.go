package rest

import (
	"expenseTracker/rest/handler"
	"expenseTracker/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProducts), middleware.Auth))

	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handler.GetProductByID))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(handler.CreateProduct), middleware.Auth))

	mux.Handle("GET /users", manager.With(http.HandlerFunc(handler.GetUsers)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handler.CreateUser), middleware.Auth))

	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handler.Login)))

}
