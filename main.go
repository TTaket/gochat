package main

import (
	"github.com/TTaket/gochat/router"
	"github.com/TTaket/gochat/utils"
)

// @title GoChat API
// @version 1.0
// @description This is a sample server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	utils.InitConfig()

	//tests.TestGorm()
	//tests.TestRedis()

	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
