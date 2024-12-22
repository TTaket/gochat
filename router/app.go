package router

import (
	docs "github.com/TTaket/gochat/docs"
	"github.com/TTaket/gochat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/TTaket/gochat/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.GetPing)
	r.GET("/index", service.GetIndex)
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)

	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/getUserList", service.GetUserList)
		admin.GET("/createUser", service.CreateUser)
		admin.GET("/getUserByID", service.GetUserByID)
		admin.GET("/deleteUserByID", service.DeleteUserByID)
		admin.POST("/updateUser", service.UpdateUser)
	}

	user := r.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	{

	}
	return r
}
