package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		fmt.Println(err)
	}
	height := ctx.PostForm("height")
	heightInt, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println(err)
	}
	price := ctx.PostForm("price")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println(err)
	}
	model.CreateProductionModel(productionCode, category, company, widthInt, heightInt, priceInt)
	ctx.Redirect(302, "/production/toppage")
}
