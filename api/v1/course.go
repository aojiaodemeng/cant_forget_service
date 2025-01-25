package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加分类
func AddCourse(c *gin.Context) {
	var data model.Course
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCourseExist(data.Title)
	if code == status_code.SUCCESS {
		model.CreateCourse(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCourseList(c *gin.Context) {
	data := model.GetCourseList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 编辑分类
func EditCourse(c *gin.Context) {
	var data model.JSONBindCourse
	id := c.Param("id")
	//id, _ := strconv.Atoi(c.Param("id"))
	//c.ShouldBindJSON(&data)
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf(err.Error())
	}

	model.EditCourse(id, &data)
	//var maps = make(map[string]interface{})
	//model.EditCourse(id, maps)
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	code = model.DeleteCourse(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}

// 获取详情
func GetCourseDetail(c *gin.Context) {
	recordId := c.Param("recordId")
	record, err := model.GetCourseDetail(recordId)

	if err != nil {
		code = status_code.ERROR
	} else {
		code = status_code.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    record,
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}
