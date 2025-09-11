package main

import "expenseTracker/cmd"

func main() {
	cmd.Serve()

	// token, _ := utils.CreateJWT("secret", utils.Payload{
	// 	Sub:         "1234567890",
	// 	Name:        "John Doe",
	// 	Email:       "4gCwI@example.com",
	// 	IsShopOwner: true,
	// })

	// fmt.Println(token)
}
