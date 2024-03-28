package main

import (
	"fmt"
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

func main() {
	mux := http.NewServeMux()

	// Register handlers for specific routes
	mux.HandleFunc("GET /message", getProductHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
