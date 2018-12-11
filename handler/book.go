package handler

import (
	"golang-web-starter/models"
	"golang-web-starter/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	books := models.GetAllMessages()
	utility.Render(
		c,
		gin.H{
			"title":   "Hi",
			"payload": books,
		},
		"index.html",
	)
}

func GetBook(c *gin.Context) {
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		book := models.GetBookByID(bookID)
		utility.Render(
			c,
			gin.H{
				"title":   "Book",
				"payload": book,
			},
			"book.html",
		)

	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}

}

func SaveBook(c *gin.Context) {
	var book models.Book
	if err := c.Bind(&book); err == nil {
		book := models.SaveBook(book)
		utility.Render(
			c,
			gin.H{
				"title":   "Save",
				"payload": book,
			},
			"success.html",
		)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
