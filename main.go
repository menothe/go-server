package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Seller      string    `json:"seller"`
	Category    string    `json:"category"`
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	product := &Product{
		Id:          uuid.New(),
		Description: "Old Leather Boots, size 11, black",
		Price:       110.34,
		Seller:      "sellerRuss",
		Category:    "Shoes",
	}
	fmt.Println(*product)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct = &Product{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Process the request body based on content type
	// (demonstrated with JSON data processing)
	if err := json.Unmarshal(body, &newProduct); err != nil {
		// Handle invalid JSON format
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	newProduct.Id = uuid.New()
	fmt.Println(*newProduct)
}

func main() {
	mux := http.NewServeMux()

	// Register handlers for specific routes
	mux.HandleFunc("GET /product", getProductHandler)
	mux.HandleFunc("POST /product", createProductHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

/*
curl -X POST localhost:8080/product \
	 -H "Content-Type: application/json" \
	 -d '{"description": "Harry Potter and the Sorcerers Stone", "price": 19.99, "seller": "sellerBob", "category": "books"}'
*/
