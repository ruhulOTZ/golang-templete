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
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	port := fmt.Sprintf(":%d", server.cnf.HttpPort)

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
