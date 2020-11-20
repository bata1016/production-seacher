package employee

import (
	"fmt"
	"net/http"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Controller EmployeeControllerを指す
type Controller struct {
	Db *gorm.DB
}

// IndexEmployeeはEmployeeのindexアクション
func (c Controller) IndexEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

// CreateEmployeeはEmployeeのcreateアクション
func (c Controller) CreateEmployee(ctx *gin.Context) {
	var model model.EmployeeModel
	name := ctx.PostForm("name")
	employeeCode := ctx.PostForm("employeeCode")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	pointer, err := model.CreateModel(ctx)

	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		ctx.JSON(201, pointer)
		ctx.Redirect(302, "/production/index")
	}
}
