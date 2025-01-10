package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加分类
func AddColumn(c *gin.Context) {
	var data model.Column
	_ = c.ShouldBindJSON(&data)
	model.CreateColumn(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

// 查询列表
func GetColumnList(c *gin.Context) {
	data := model.GetColumnList()
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}
