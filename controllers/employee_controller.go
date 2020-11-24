package controllers

import (
	"net/http"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
)

// Controller EmployeeControllerを指す
// type EmployeeController struct {
// 	Db *gorm.DB
// }

// IndexEmployeeはEmployeeのindexアクション
func IndexEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

// CreateEmployeeはEmployeeのcreateアクション
func CreateEmployee(ctx *gin.Context) {
	var model model.EmployeeModel
	name := ctx.PostForm("name")
	employeeCode := ctx.PostForm("employeeCode")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	model.CreateModel(name, employeeCode, email, password)
	// ctx.Redirect(302, "/production/toppage")
	// if err != nil {
	// 	ctx.AbortWithStatus(400)
	// 	fmt.Println(err)
	// } else {
	// 	ctx.JSON(201, pointer)
	// 	ctx.Redirect(302, "/production/index")
	// }
}
