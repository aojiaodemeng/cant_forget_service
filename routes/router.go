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
		auth.POST("file/category/add", v1.AddFileCategory)
		auth.GET("file/category/list", v1.GetFileCategories)
		auth.PUT("file/category/:id", v1.EditFileCategory)
		auth.DELETE("file/category/:id", v1.DeleteFileCategory)
	}

	router := r.Group("api/v1")
	{
		router.POST("user/register", v1.RegisterUser)
		router.POST("login", v1.Login)

		router.POST("course/category/add", v1.AddCourseCategory)
		router.GET("course/category/list", v1.GetCourseCategories)
		router.PUT("course/category/:id", v1.EditCourseCategory)
		router.DELETE("course/category/:id", v1.DeleteCourseCategory)

	}

	r.Run(utils.HttpPort)
}
