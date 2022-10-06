package routes

import (
	"1.19/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/bookstore/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/bookstore/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/bookstore/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/bookstore/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/bookstore/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
