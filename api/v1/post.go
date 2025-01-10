package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddPost(c *gin.Context) {
	var r model.CreatePostReqParams
	_ = c.ShouldBindJSON(&r)

	var courseCategoryList []model.CourseCategory
	for _, v := range r.CategoryIds {
		courseCategoryList = append(courseCategoryList, model.CourseCategory{
			Id: v,
		})
	}

	var columns []model.Column
	for _, v := range r.ColumnIds {
		columns = append(columns, model.Column{
			Id: v,
		})
	}

	fmt.Printf(strconv.Itoa(int(r.PostTypeId)))

	data := &model.Post{Title: r.Title, Content: r.Content, Img: r.Img, CourseCategories: courseCategoryList, Columns: columns}
	//code = model.CreatePost(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}
