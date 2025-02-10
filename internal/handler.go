package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fthilov/devops-lecture-project/pkg/auth"
	"github.com/fthilov/devops-lecture-project/pkg/products"
	"github.com/gorilla/mux"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := r.FormValue("username")
	password := r.FormValue("password")
	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := auth.CreateToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error generating the token"}`))
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid credentials"}`))
	}
}
func AuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// In this simple example, we'll just return a success message
	w.Write([]byte(`{"message": "Logout successful"}`))
}
func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(products.Products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Write(response)
}
func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID is missing"}`))
		return
	}
	id, err := strconv.Atoi(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID has wrong format"}`))
		return
	}
	product := products.FindProductByID(products.Products, id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}
	response, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Write(response)
}
func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Missing authorization header"}`))
		return
	}
	if auth.VerifyToken(token) {
		// In this simple example, we'll just return a success message
		w.Write([]byte(`{"message": "Order placed successfully"}`))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid token"}`))
	}
}
