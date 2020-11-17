package main

import (
	"github.com/bata1016/production-seacher/db"
	employee "github.com/bata1016/production-seacher/server"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()
	router.Static("/assets/", "./assets")
	employee.EmployeeRouter(router)
	router.Run()
}
