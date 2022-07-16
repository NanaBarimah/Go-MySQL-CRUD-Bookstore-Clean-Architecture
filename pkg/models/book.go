package models

import (
	"github.com/NanaBarimah/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Struct basically are based on models
// And models gives us a structure to help us store something in the database
type Book struct {
	gorm.Model
	Name        string `gorm: ""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialze the database
func init() {
	config.Connect()
	db = config.GetDB()     // The db returned from this GetDB function is stored in the db variable above
	db.AutoMigrate(&Book{}) // Auto migrate this with an empty book structure
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book { // using a slice because we want to return a slice or list of books in the database
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
