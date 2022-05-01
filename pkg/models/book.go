package models

import(
	"github.com/jinzhu/gorm"
	"github.com/Quddus1916/Book-store_golang/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"Name"`
	Author string `json:"Author"`
	Publication string `json:"Publication"`

}

func init(){
     
}