package router

import (
	"github.com/TTaket/gochat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetPing)
	r.GET("/index", service.GetIndex)

	return r
}
