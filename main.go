package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
}

func main() {
	mux := http.NewServeMux()

	// Register handlers for specific routes
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/info", infoHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
