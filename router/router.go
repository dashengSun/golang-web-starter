package router

import (
	"github.com/gin-gonic/gin"
	"golang-web-starter/handler"
)

// Load loads the routes
func Load() *gin.Engine {
	router := gin.Default()

	router.GET("/", handler.ShowIndexPage)
	router.GET("/book/:book_id", handler.GetBook)
	router.POST("/book", handler.SaveBook)

	return router
}
