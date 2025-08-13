package cmd

import (
	"expenseTracker/global_router"
	"expenseTracker/handler"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handler.GetProductByID))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :4000")
	err := http.ListenAndServe(":4000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
