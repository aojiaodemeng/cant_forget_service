package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加分类
func AddFileCategory(c *gin.Context) {
	var data model.FileCategory
	_ = c.ShouldBindJSON(&data)
	code = model.CheckFileCategoryExist(data.ID)
	if code == status_code.SUCCESS {
		model.CreateFileCategory(&data)
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
func GetFileCategories(c *gin.Context) {
	data := model.GetFileCategories()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 编辑分类
func EditFileCategory(c *gin.Context) {
	var data model.FileCategory
	id := c.Param("id")
	c.ShouldBindJSON(&data)
	model.EditFileCategory(id, &data)
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}

// 删除分类
func DeleteFileCategory(c *gin.Context) {
	id := c.Param("id")
	code = model.DeleteFileCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}
