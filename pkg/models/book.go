package models

import(
	"github.com/jinzhu/gorm"
	"github.com/Quddus1916/Book-store_golang/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"Name"`
	Author string `json:"Author"`
	Publication string `json:"Publication"`

}

func init(){
	config.Connect()
	
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func  CreateBook(b *Book) *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func  GetBookById(id int64) (*Book ,*gorm.DB){
	var BookById Book
	db:= db.Where("ID=?",id).Find(&BookById)
	return &BookById ,db
 
}

func DeleteBook(id int64) Book{
	var Book Book
	 db.Where("ID=?",id).Delete(Book)
	return Book

}