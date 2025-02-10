package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fthilov/devops-lecture-project/pkg/products"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", internal.authLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.authLogoutHandler).Methods("POST")
	// Product Service
	router.HandleFunc("/products", internal.productListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", internal.productDetailHandler).Methods("GET")
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", internal.checkoutPlaceOrderHandler).Methods("POST")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func findProductByID(products []products.Product, id int) *products.Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
