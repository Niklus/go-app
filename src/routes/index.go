package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	homeRouter(router)
	userRouter(router)
	return router
}
