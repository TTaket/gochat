package router

import (
	docs "github.com/TTaket/gochat/docs"
	"github.com/TTaket/gochat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.GetPing)
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/getUserByID", service.GetUserByID)
	r.GET("/user/deleteUserByID", service.DeleteUserByID)
	r.POST("/user/updateUser", service.UpdateUser)

	return r
}
