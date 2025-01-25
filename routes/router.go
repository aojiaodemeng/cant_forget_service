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

		router.POST("course/add", v1.AddCourse)
		router.GET("course/list", v1.GetCourseList)
		router.PUT("course/:id", v1.EditCourse)
		router.GET("course/detail/:recordId", v1.GetCourseDetail)
		//router.GET("course/detail/:recordId/linked_records", v1.GetCourseCategoryLinkedRecords)

		router.POST("post/add", v1.AddPost)
		router.GET("post/list", v1.GetPostList)
		router.PUT("post/:id", v1.EditCourse)
		router.DELETE("course/:id", v1.DeleteCourse)

		router.POST("column/add", v1.AddColumn)
		router.GET("column/list", v1.GetColumnList)

		router.POST("task/add", v1.AddTask)
		router.GET("task/list", v1.GetTaskList)
		router.POST("task/addColumn/:columnId", v1.AddColumnToTask)
		router.GET("task/todayTodoList", v1.GetTodayTodoTaskList)
		//router.GET("task/planedList", v1.GetTodayTodoTaskList)
		router.GET("task/todolist", v1.GetTodoList)
	}

	r.Run(utils.HttpPort)
}
