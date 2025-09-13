package cmd

import (
	"expenseTracker/config"
	"expenseTracker/rest"
	"expenseTracker/rest/handler/product"
	"expenseTracker/rest/handler/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	rest.NewServer(productHandler, userHandler).Start(cnf) // start the server
}
