package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Seller      string    `json:"seller"`
	Category    string    `json:"category"`
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	product := &Product{
		Id:          uuid.New(),
		Description: "Old Leather Boots, size 11, black",
		Price:       110.34,
		Seller:      "sellerRuss",
		Category:    "Shoes",
	}
	fmt.Println(*product)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct = &Product{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Process the request body based on content type
	// (demonstrated with JSON data processing)
	if err := json.Unmarshal(body, newProduct); err != nil {
		// Handle invalid JSON format
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	newProduct.Id = uuid.New()

	err = saveData(newProduct)
	if err != nil {
		fmt.Println("Unable to create product:", err)
		return
	}

	fmt.Println("product created: ", *newProduct)
}

func FetchAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product

	err := readUsersFromFile(&products, "./products.json")
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}
	fmt.Println("all products: ", products)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var products []Product

	err := readUsersFromFile(&products, "./products.json")
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	for index, product := range products {
		if product.Id.String() == id {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	fmt.Printf("deleted product with id %s \n", id)

	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	err = os.WriteFile("./products.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return
	}
}

// Helper function to read existing data (modify as needed)
func readUsersFromFile(products *[]Product, filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil && os.IsNotExist(err) {
		// File doesn't exist, initialize empty slice
		*products = []Product{}
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(b, products)
}

func saveData(product *Product) error {
	var products []Product

	err := readUsersFromFile(&products, "./products.json")
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return err
	}
	products = append(products, *product)

	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return err
	}
	err = os.WriteFile("./products.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return err
	}
	fmt.Println("Data persisted to products.json")
	return nil
}
