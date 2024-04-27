package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyphenhypen/go-bookstore-backend/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
