package main

import (
	"bookStore/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("server is listening on port: 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
