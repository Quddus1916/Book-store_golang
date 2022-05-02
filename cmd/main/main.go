package main

import (
	"log"
	"net/http"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"github.com/Quddus1916/Book-store_golang/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080",r))
	
}