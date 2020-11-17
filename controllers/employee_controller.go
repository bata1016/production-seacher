package employee

import (
	"fmt"
	"net/http"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c Controller) IndexEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
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
