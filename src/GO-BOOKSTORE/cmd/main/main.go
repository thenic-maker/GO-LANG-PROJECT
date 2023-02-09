package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm/dialects/mysql"
	"GO-BOOKSTORE/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("locathost:9010", r))
}
