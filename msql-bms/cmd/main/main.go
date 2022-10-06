package main

import (
	"fmt"
	"log"
	"net/http"

	"1.19/pkg/routes"
	"github.com/gorilla/mux"
	// "gorm.io/gorm/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	fmt.Println("Starting the application...")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
