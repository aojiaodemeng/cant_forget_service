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
		router.POST("user/register", v1.RegisterUser)
		router.POST("login", v1.Login)

		router.POST("course/category/add", v1.AddCourseCategory)
		router.GET("course/category/list", v1.GetCourseCategoryList)
		router.PUT("course/category/:id", v1.EditCourseCategory)
		router.DELETE("course/category/:id", v1.DeleteCourseCategory)

		router.POST("post/add", v1.AddPost)
		router.GET("post/list", v1.GetCourseCategoryList)
		router.PUT("post/:id", v1.EditCourseCategory)
		router.DELETE("course/:id", v1.DeleteCourseCategory)

		router.POST("post/type/add", v1.AddPostType)
		router.GET("post/type/list", v1.GetPostTypeList)
		router.PUT("post/type/:id", v1.EditPostType)
		router.DELETE("post/type/:id", v1.DeletePostType)

		router.POST("column/add", v1.AddColumn)
		router.GET("column/list", v1.GetColumnList)
	}

	r.Run(utils.HttpPort)
}
