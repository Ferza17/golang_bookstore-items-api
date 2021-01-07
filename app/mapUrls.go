package app

import (
	"github.com/Ferza17/golang_bookstore-items-api/controllers"
	"net/http"
)

func Urls() {
	// Create Item URL
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	//Get By Id
	r.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	//Search
	r.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
