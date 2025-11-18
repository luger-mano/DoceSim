package controller

import (
	"DoceSim/database"
	"DoceSim/model"

	"github.com/gin-gonic/gin"
)

func ProductsCreate(ctx *gin.Context) {
	product := model.Product{Name: "Brigadeiro", Description: "Brigadeiro com confetes", Value: 9.5}

	response := database.DB.Create(&product)

	if response.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"product": product,
	})
}
