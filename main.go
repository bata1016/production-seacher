package main

import (
	"github.com/bata1016/production-seacher/controllers"
	"github.com/bata1016/production-seacher/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("/assets/", "./assets")
	router.GET("/", controllers.IndexEmployee)
	router.GET("/login", controllers.LoginEmployee)
	router.POST("/newemployee", controllers.CreateEmployee)

	production := router.Group("/production")
	{
		production.GET("/toppage", controllers.ProductionIndex)
	}
	router.Run()
}
