package main

import (
	"golang-web-starter/router"
)

func main() {
	router := router.Load()
	router.LoadHTMLGlob("templates/*")
	router.Run() // listen and serve on 0.0.0.0:8080
}
