package main

import (
	"alexandria.backend/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "alexandria.backend/docs"
)

func main() {
	routes := gin.Default()

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.GET("/books/:isbn", controllers.GetBookByIsbn)

	routes.Run()
}
