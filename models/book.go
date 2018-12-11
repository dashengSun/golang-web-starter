package models

import (
	"golang-web-starter/utility"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `form:"title" json:"title" gorm:"type:varchar(100);unique"`
	Author string `form:"author" json:"author" gorm:"type:varchar(100)"`
}

func GetAllBooks() []Book {
	var books []Book
	utility.DB().Find(&books)
	return books
}

func GetBookByID(id int) Book {
	var book Book
	utility.DB().First(&book, id)
	return book
}

func SaveBook(book Book) Book {
	utility.DB().Create(&book)
	return book;
}
