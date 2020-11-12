package main

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/server"
)

func main() {
	db.Init()
	server.Router()
}
