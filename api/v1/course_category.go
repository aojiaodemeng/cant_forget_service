package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加分类
func AddCourseCategory(c *gin.Context) {
	var data model.CourseCategory
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCourseCategoryExist(data.ID)
	if code == status_code.SUCCESS {
		model.CreateCourseCategory(&data)
	}
	if code == status_code.ERROR_FILE_CATEGORY_USED {
		code = status_code.ERROR_FILE_CATEGORY_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCourseCategories(c *gin.Context) {
	data := model.GetCourseCategories()
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
	id := c.Param("id")
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
