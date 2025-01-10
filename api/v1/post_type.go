package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddPostType(c *gin.Context) {
	var data model.PostType
	_ = c.ShouldBindJSON(&data)
	code = model.CheckPostTypeExist(data.Name)
	if code == status_code.SUCCESS {
		model.CreatePostType(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询分类列表
func GetPostTypeList(c *gin.Context) {
	data := model.GetFileCategories()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 编辑分类
func EditPostType(c *gin.Context) {
	var data model.PostType
	//id := c.Param("id")
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	model.EditPostType(id, &data)
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}

// 删除分类
func DeletePostType(c *gin.Context) {
	id := c.Param("id")
	code = model.DeletePostType(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": status_code.GetErrMsg(code),
	})
}
