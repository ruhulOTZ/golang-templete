package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	sendData(w, productList, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func handleCorsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "Bearer")
		w.Header().Set("Content-Type", "application/json")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleCors(w, r)
		next.ServeHTTP(w, r)
	})
}

func handleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "Bearer")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	sendData(w, productList, http.StatusOK)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server running on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func init() {
	productList = []Product{
		{
			ID:          1,
			Title:       "Apple MacBook Air",
			Description: "Laptop",
			Price:       999.99,
			ImgUrl:      "https://picsum.photos/200/300",
		},
		{
			ID:          2,
			Title:       "Dell XPS 13",
			Description: "Laptop",
			Price:       1299.99,
			ImgUrl:      "https://picsum.photos/200/301",
		},
		{
			ID:          3,
			Title:       "Canon EOS 80D",
			Description: "Camera",
			Price:       1200.00,
			ImgUrl:      "https://picsum.photos/200/302",
		},
	}
}
