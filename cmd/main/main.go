package main

import (
	"log"
	"net/http"

	_ "gihub.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/vivekhiwarkar/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
