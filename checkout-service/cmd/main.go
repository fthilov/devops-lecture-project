package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fthilov/devops-lecture-project/checkout-service/internal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", internal.CheckoutPlaceOrderHandler).Methods("POST")
	port := 8081
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
