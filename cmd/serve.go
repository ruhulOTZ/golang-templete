package cmd

import (
	"expenseTracker/config"
	"expenseTracker/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf) // start the server
}
