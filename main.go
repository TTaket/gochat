package main

import (
	"github.com/TTaket/gochat/router"
	"github.com/TTaket/gochat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	//tests.TestGorm()

	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
