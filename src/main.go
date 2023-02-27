package main

import (
	"fmt"
	"go-app/routes"
	"net/http"
)

func main() {
	
	router := routes.Router()
	fs := http.FileServer(http.Dir("static"))
	port := ":8080"
	
	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
