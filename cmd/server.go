package cmd

import (
	"expenseTracker/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on :4000")
	err := http.ListenAndServe(":4000", wrapperMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
