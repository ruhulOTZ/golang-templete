package product

import (
	"encoding/json"
	"expenseTracker/database"
	"expenseTracker/rest/middleware"
	"expenseTracker/utils"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.GetProducts(), http.StatusOK)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	product := database.GetProductByID(id)

	if err != nil {
		utils.SendData(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if product == nil {
		utils.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	response := database.AddProduct(&newProduct)

	utils.SendData(w, response, 201)
}
