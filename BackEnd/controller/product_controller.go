package controller

import (
	"DoceSim/database"
	"DoceSim/model"

	"github.com/gin-gonic/gin"
)

func ProductsCreate(ctx *gin.Context) {

	var body struct {
		Name        string
		Description string
		Value       float64
	}

	ctx.Bind(&body)

	product := model.Product{Name: body.Name, Description: body.Description, Value: body.Value}
	response := database.DB.Create(&product)

	if response.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"product": product,
	})
}

func ProductsShow(ctx *gin.Context) {
	var products []model.Product
	database.DB.Find(&products)

	ctx.JSON(200, gin.H{
		"products": products,
	})
}

func ProductsIndex(ctx *gin.Context) {
	id := ctx.Param("id")

	var product model.Product
	database.DB.First(&product, id)

	ctx.JSON(200, gin.H{
		"product": product,
	})
}

func ProductsUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Name        string
		Description string
		Value       float64
	}

	ctx.Bind(&body)

	var product model.Product
	database.DB.First(&product, id)

	database.DB.Model(&product).Updates(model.Product{Name: body.Name, Description: body.Description, Value: body.Value})

	ctx.JSON(200, gin.H{
		"product": product,
	})
}
func ProductsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var product model.Product
	database.DB.First(&product, id)

	database.DB.Delete(&product)

	ctx.JSON(200, gin.H{
		"product": product,
	})
}
