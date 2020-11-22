package main

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/server"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()
	router.Static("/assets/", "./assets")
	server.EmployeeRouter(router)
	server.ProductionRouter(router)
	router.Run()
}
