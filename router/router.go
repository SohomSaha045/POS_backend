package router

import (
	"github.com/SohomSaha045/POS_backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/item",controller.ItemController).Methods("GET")
	return router
}