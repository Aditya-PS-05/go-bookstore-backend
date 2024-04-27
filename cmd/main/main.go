package main

import (
	"log"
	"net/http"

	"github.com/Hyphenhypen/go-bookstore-backend/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
