package models

import (
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB;


type Book struct {
	gorm.Model
	Id 		string `json:"id"`
	Title 	string `json:"title"`
	Author	string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	db = config.Connect();
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(b);
	return b;
}

func GetAllBooks() []Book{
	// here we have the address the &books
	var books []Book;
	result := db.Find(&books)
	if result.Error != nil {
		return nil
	}

	return books;
}

// if we are retuning the *Book means we have to return the address
// *Book means you are returning a pointer to a Book.
// &book is getting the address of the book variable, which is exactly what a pointer expects.
func GetBookById(Id int64) (*Book, *gorm.DB){
	// here we have the address the &books
	var book Book;
	// func (db *gorm.DB) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	db.Where("Id ?= ", Id).Find(&book);
	return &book, db;
}

func DeleteBookById(Id int64) any {
	// here we have the address the &books
	var book Book;
	db.Where("Id ?= ", Id).Delete(book);
	return book
}

