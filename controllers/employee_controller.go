package controllers

import (
	"log"
	"net/http"

	model "github.com/bata1016/production-seacher/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Controller EmployeeControllerを指す
// type EmployeeController struct {
// 	Db *gorm.DB
// }

// IndexEmployeeはEmployeeのindexアクション
func IndexEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

// LoginEmployeeはLoginページを返却するアクション
func LoginEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// CreateEmployeeはEmployeeのcreateアクション
func CreateEmployee(ctx *gin.Context) {
	var model model.EmployeeModel
	name := ctx.PostForm("name")
	employeeCode := ctx.PostForm("employeeCode")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	model.CreateModel(name, employeeCode, email, password)
	ctx.Redirect(302, "/production/toppage")
}

func SessionCheckEmployee(ctx *gin.Context) {
	var model model.EmployeeModel
	employeeCode := ctx.PostForm("employeeCode")
	// password := ctx.PostForm("password")
	dbPassWord := model.GetLoginEmployee(employeeCode).Password
	formPassword := ctx.PostForm("password")
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassWord), []byte(formPassword)); err != nil {
		log.Println("ログインできませんでした")
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{})
	} else {
		log.Println("ログインできました")
		ctx.Redirect(302, "/production/toppage")
	}
	// var model model.EmployeeModel
	// employeeCode := ctx.PostForm("employeeCode")
	// password := ctx.PostForm("password")

}
