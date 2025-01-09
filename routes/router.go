package routes

import (
	v1 "cant_forget/api/v1"
	"cant_forget/middleware"
	"cant_forget/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		auth.GET("user/editUserInfo", v1.EditUserInfo)
	}

	router := r.Group("api/v1")
	{
		auth.POST("user/register", v1.RegisterUser)
		router.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
