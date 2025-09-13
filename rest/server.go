package rest

import (
	"expenseTracker/config"
	"expenseTracker/rest/handler/product"
	"expenseTracker/rest/handler/user"
	"expenseTracker/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	port := fmt.Sprintf(":%d", cnf.HttpPort)

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)

	// initRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(port, wrapperMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
