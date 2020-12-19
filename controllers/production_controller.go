package controllers

import (
	"fmt"
	"net/http"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
)

// type ProductionController struct {
// 	gorm.Model
// }

func ProductionIndex(ctx *gin.Context) {
	var production model.ProductionModel
	productions, err := production.GetAll()
	if err != nil {
		fmt.Println(err)
	} else {
		ctx.HTML(http.StatusOK, "toppage.html", gin.H{
			"productions": productions,
		})
	}
}

func ProductionNew(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addcost.html", gin.H{})
}

func ProductionCreate(ctx *gin.Context) {
	var model model.ProductionModel
	productionCode := ctx.PostForm("productionCode")
	category := ctx.PostForm("category")
	company := ctx.PostForm("company")
	width := ctx.PostForm("width")
	height := ctx.PostForm("height")
	price := ctx.PostForm("price")
	errorMessages := model.CreateProductionModel(productionCode, category, company, width, height, price)
	if errorMessages != nil {
		ctx.HTML(http.StatusBadRequest, "addcost.html", gin.H{
			"errorMessages": errorMessages,
		})
	} else {
		ctx.Redirect(302, "/production/toppage")
	}
}
