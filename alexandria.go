package main

import (
	"alexandria.backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	routes.GET("/books/:isbn", controllers.GetBookByIsbn)

	routes.Run()
}
