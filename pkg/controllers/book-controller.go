package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/Quddus1916/Book-store_golang/pkg/utils"
	"github.com/Quddus1916/Book-store_golang/pkg/models"
)
var NewBook  models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	
}



func GetBookById(w http.ResponseWriter, r *http.Request) {

	
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {

	
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	
}