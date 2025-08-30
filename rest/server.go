package rest

import (
	"expenseTracker/config"
	"expenseTracker/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

func Start(cnf config.Config) {
	port := fmt.Sprintf(":%d", cnf.HttpPort)

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(port, wrapperMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
