package main

import (
	"github.com/Guhfrontend/gopportunities/controller"
	"github.com/Guhfrontend/gopportunities/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductUsecase := usecase.NewProductUsecase()

	ProductController := controller.NewProductController(*ProductUsecase)

	server.GET("/products", ProductController.GetProducts)
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8080")
}
