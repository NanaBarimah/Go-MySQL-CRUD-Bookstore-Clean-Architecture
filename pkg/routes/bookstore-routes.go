package routes

import (
	"github.com/NanaBarimah/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// Create a function
// This function will have all the routes
// The routes will help get the control to the controller where
// it will contain all the logic
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
