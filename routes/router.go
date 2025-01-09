package routes

import (
	v1 "cant_forget/api/v1"
	"cant_forget/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.POST("user/register", v1.RegisterUser)
		router.GET("user/editUserInfo", v1.EditUserInfo)
		router.GET("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
