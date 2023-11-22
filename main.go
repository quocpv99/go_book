package main

import (
	"fmt"
	"go_bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
