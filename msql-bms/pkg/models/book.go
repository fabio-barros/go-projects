package models

import (
	"1.19/pkg/config"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []*Book {
	var books []*Book
	db.Find(&books)
	return books
}

func GetBook(id int) *Book {
	var book Book
	db.First(&book, id)
	return &book
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func (b *Book) DeleteBook() *Book {
	db.Delete(&b)
	return b
}
