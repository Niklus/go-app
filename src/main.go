package main

import (
	"fmt"
	"go-app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	port := ":8080"

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
	
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
