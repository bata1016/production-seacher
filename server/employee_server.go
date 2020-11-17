package employee

import (
	employee "github.com/bata1016/production-seacher/controllers"
	"github.com/gin-gonic/gin"
)

func EmployeeRouter(route *gin.Engine) {
	route.LoadHTMLGlob("templates/employee/*.html")
	{
		employeeCtrl := employee.Controller{}
		route.GET("", employeeCtrl.IndexEmployee)
		route.POST("", employeeCtrl.CreateEmployee)
	}
}
