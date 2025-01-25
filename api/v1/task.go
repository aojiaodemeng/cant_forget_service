package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加文章
func AddTask(c *gin.Context) {
	var r model.CreateTaskReqParams
	_ = c.ShouldBindJSON(&r)
	data := &model.CreateTaskReqParams{N: r.N, Efactor: r.Efactor, Users: r.Users, Posts: r.Posts, Interval: r.Interval}
	code = model.CreateTask(data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    r,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询列表
func GetTaskList(c *gin.Context) {
	//columnUuid, _ := strconv.Atoi(c.Query("columnUuid"))
	//columnUuid := c.Query("columnUuid")
	data := model.GetTaskList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 专栏加入学习计划
func AddColumnToTask(c *gin.Context) {
	columnId := c.Param("columnId")
	code := model.AddColumnToTask(columnId)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		//"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 获取当天任务列表（卡片）
func GetTodayTodoTaskList(c *gin.Context) {
	data := model.GetTodayTodoTaskList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 获取当天任务列表
func GetTodoList(c *gin.Context) {
	data := model.GetTodoList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}
