package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

)

var (
	db *gorm.DB
)

func connect() {
	d,err:= gorm.Open("mysql","root:quddus_1916@tcp(127.0.0.1:3306)/book_management?charset=utf8&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	db = d
}

func test(){
	log.Panic("test")
}

func GetDB() *gorm.DB{
	return db
}