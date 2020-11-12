package employee

import (
	"fmt"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c Controller) IndexEmployee(ctx *gin.Context) {
	var employee model.EmployeeModel
	pointer, err := employee.GetAll()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, pointer)
	}
}

func (c Controller) CreateEmployee(ctx *gin.Context) {
	var model model.EmployeeModel
	pointer, err := model.CreateModel(ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(201, pointer)
	}
}
