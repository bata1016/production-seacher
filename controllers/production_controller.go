package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type ProductionController struct {
// 	gorm.Model
// }

func ProductionIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
