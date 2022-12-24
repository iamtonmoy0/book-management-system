package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamtonmoy0/book-management-system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/", controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("Get")
}
