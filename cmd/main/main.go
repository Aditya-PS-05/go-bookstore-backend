package main

import (
	"fmt"
	"net/http"

	"github.com/Aditya-PS-05/go-bookstore-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Muhh")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8000", r)
	fmt.Println("The server is running on port 8000")

}
