package server

import (
	"github.com/bata1016/production-seacher/controllers"
	"github.com/gin-gonic/gin"
)

func EmployeeRouter(route *gin.Engine) {
	route.LoadHTMLGlob("templates/employee/*.html")
	{
		employeeCtrl := controllers.EmployeeController{}
		route.GET("", employeeCtrl.IndexEmployee)
		route.POST("", employeeCtrl.CreateEmployee)
	}
}
