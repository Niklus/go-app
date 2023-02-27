package routes

import (
	"go-app/handlers"

	"github.com/gorilla/mux"
)

func homeRouter(r*mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
}
