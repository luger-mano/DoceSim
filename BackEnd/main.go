package main

import (
	"DoceSim/controller"
	. "DoceSim/database"
	"DoceSim/model"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	ConnectDB()

	DB.AutoMigrate(&model.Product{})

	server.POST("/products", controller.ProductsCreate)
	server.PUT("/products/:id", controller.ProductsUpdate)
	server.DELETE("/products/:id", controller.ProductsDelete)
	server.GET("/products", controller.ProductsShow)
	server.GET("/products/:id", controller.ProductsIndex)

	server.Run(":8000")
}
