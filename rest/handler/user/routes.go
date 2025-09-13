package user

import (
	"expenseTracker/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetUsers)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser), middleware.Auth))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))

}
