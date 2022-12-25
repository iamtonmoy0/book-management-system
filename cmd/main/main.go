package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/iamtonmoy0/book-management-system/pkg/routes"
	_ "github.com/jinzhu/gorm/dailects/mysql"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
