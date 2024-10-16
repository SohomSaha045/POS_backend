package router

import (
	"github.com/SohomSaha045/POS_backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/item",controller.GetAllItemController).Methods("GET")
	router.HandleFunc("/item",controller.AddItemController).Methods("POST")
	return router
}