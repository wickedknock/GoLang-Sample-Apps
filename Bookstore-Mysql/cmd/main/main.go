package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wickedknock/Bookstore-Mysql/pkg/routes"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
