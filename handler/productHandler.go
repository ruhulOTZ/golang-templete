package handler

import (
	"encoding/json"
	"expenseTracker/database"
	"expenseTracker/utils"
	"fmt"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductList, http.StatusOK)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.SendData(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		fmt.Println(product.ID == id)
		if product.ID == id {
			fmt.Println(id)
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}

	utils.SendData(w, "Product not found", http.StatusNotFound)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)
}
