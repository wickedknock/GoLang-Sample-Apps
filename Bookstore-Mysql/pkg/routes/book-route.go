package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/wickedknock/Bookstore-Mysql/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
