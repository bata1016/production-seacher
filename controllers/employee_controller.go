package employee

import (
	"net/http"

	"github.com/bata1016/production-seacher/models/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c Controller) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	employeeCode := ctx.PostForm("employeeCode")
	email := ctx.PostForm("email")
	c.Db.Create(&entity.Employee{Name: name, EmployeeCode: employeeCode, Email: email})
	ctx.Redirect(http.StatusMovedPermanently, "/")
}
