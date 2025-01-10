package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCourseCategory(c *gin.Context) {
	var data model.CourseCategory
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCourseCategoryExist(data.ID)
	if code == status_code.SUCCESS {
		model.CreateCourseCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCourseCategoryList(c *gin.Context) {
	data := model.GetCourseCategoryList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 编辑分类
func EditCourseCategory(c *gin.Context) {
	var data model.CourseCategory
	//id := c.Param("id")
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	model.EditCourseCategory(id, &data)
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCourseCategory(c *gin.Context) {
	id := c.Param("id")
	code = model.DeleteCourseCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}
