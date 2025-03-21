package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fthilov/devops-lecture-project/auth-service/internal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", internal.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.AuthLogoutHandler).Methods("POST")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
