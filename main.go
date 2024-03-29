package main

import (
	"fmt"
	"net/http"

	"github.com/menothe/go-server/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Register handlers for specific routes
	mux.HandleFunc("GET /product", handlers.GetProductHandler)
	mux.HandleFunc("DELETE /product/{id}", handlers.DeleteProductHandler)
	mux.HandleFunc("POST /product", handlers.CreateProductHandler)
	mux.HandleFunc("GET /products", handlers.FetchAllProductsHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

/*
curl -X POST localhost:8080/product \
	 -H "Content-Type: application/json" \
	 -d '{"description": "The Matrix", "price": 9.99, "seller": "sellerTom", "category": "movies"}'

curl -X POST localhost:8080/product \
	 -H "Content-Type: application/json" \
	 -d '{"description": "Harry Potter and the Sorcerers Stone", "price": 19.99, "seller": "sellerRuss", "category": "books"}'
*/
