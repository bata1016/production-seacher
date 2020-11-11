package server

import (
	employee "github.com/bata1016/production-seacher/controllers"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	router := gin.Default()
	employeePath := router.Group("/employees")
	{
		employeeCtrl := employee.Controller{}
		employeePath.GET("", employeeCtrl.IndexEmployee)
		employeePath.POST("", employeeCtrl.CreateEmployee)
	}
	return router
}
