package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fthilov/devops-lecture-project/product-service/internal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Product Service
	router.HandleFunc("/products", internal.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", internal.ProductDetailHandler).Methods("GET")
	port := 8082
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
