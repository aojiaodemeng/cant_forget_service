package v1

import (
	"cant_forget/model"
	"cant_forget/utils/status_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加文章
func AddPost1(c *gin.Context) {
	var r model.CreatePostReqParams
	_ = c.ShouldBindJSON(&r)

	//var courseCategoryList []model.Course
	//for _, v := range r.CategoryIds {
	//	courseCategoryList = append(courseCategoryList, model.Course{
	//		Id: v,
	//	})
	//}

	//var columns []model.Column
	//for _, v := range r.ColumnIds {
	//	columns = append(columns, model.Column{
	//		Id: v,
	//	})
	//}

	//fmt.Printf(strconv.Itoa(int(r.PostTypeId)))

	//data := &model.Post{Title: r.Title, Content: r.Content, Img: r.Img, CourseCategories: courseCategoryList, Columns: columns}
	//code = model.CreatePost(data)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		//"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}

func AddPost(c *gin.Context) {
	var r model.CreatePostReqParams
	_ = c.ShouldBindJSON(&r)

	//var courseCategoryList []model.Course
	//for _, v := range r.CategoryIds {
	//	courseCategoryList = append(courseCategoryList, model.Course{
	//		Id: v,
	//	})
	//}

	//var columns []string
	//for _, v := range r.ColumnIds {
	//	columns = append(columns, model.Column{
	//		Id: v,
	//	})
	//}

	//var columns []model.Column
	//for _, v := range r.Columns {
	//	columns = append(columns, model.Column{
	//		Id: v,
	//	})
	//}
	data := &model.CreatePostReqParams{Title: r.Title, Content: r.Content, Type: r.Type, Img: r.Img, Columns: r.Columns, Courses: r.Courses, Efactor: r.Efactor}
	code = model.CreatePost(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})

}

// 查询列表
func GetPostList(c *gin.Context) {
	//columnUuid, _ := strconv.Atoi(c.Query("columnUuid"))
	columnUuid := c.Query("columnUuid")
	data := model.GetPostList(columnUuid)
	code = status_code.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": status_code.GetErrMsg(code),
	})
}
