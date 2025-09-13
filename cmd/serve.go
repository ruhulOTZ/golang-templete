package cmd

import (
	"expenseTracker/config"
	"expenseTracker/rest"
	"expenseTracker/rest/handler/product"
	"expenseTracker/rest/handler/user"
	"expenseTracker/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)

	rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	).Start() // start the server
}
