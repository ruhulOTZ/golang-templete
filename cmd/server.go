package cmd

import (
	"expenseTracker/global_router"
	"expenseTracker/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	// mux.Handle("GET route", middleware.Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :4000")
	err := http.ListenAndServe(":4000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
