package server

import (
	"github.com/bata1016/production-seacher/controllers"
	"github.com/gin-gonic/gin"
)

func ProductionRouter(route *gin.Engine) {
	route.LoadHTMLGlob("templates/production/*.html")
	{
		productionCtrl := controllers.ProductionController{}
		route.GET("/production/index", productionCtrl.ProductionIndex)
	}
}
