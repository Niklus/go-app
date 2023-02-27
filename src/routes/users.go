package routes

import (
	"go-app/handlers"

	"github.com/gorilla/mux"
)

func userRouter(r*mux.Router)  {
	r.HandleFunc("/users/{id}", handlers.UserHandler).Methods("GET")
}
