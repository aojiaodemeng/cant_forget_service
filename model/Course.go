package model

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"fmt"
	"github.com/mehanizm/airtable"
	"gorm.io/gorm"
)

type JSONBindCourse struct {
	Title *string `json:"title,omitempty"`
	Desc  *string `json:"desc,omitempty"`
}
type Course struct {
	gorm.Model
	Id    uint   `gorm:"not null;unique;primary_key;comment:课程类型ID;size:90"`
	Title string `gorm:"type:varchar(20);not null"`
	Desc  string `gorm:"type:varchar(200)"`
}

type CourseLinkedData struct {
	Columns []*airtable.Record
	Posts   []*airtable.Record
}

// 查询分类是否存在
func CheckCourseExist(title string) (code int) {
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	filterFormula := fmt.Sprintf("AND({title}='%s')", title)
	records, _ := table.GetRecords().
		WithFilterFormula(filterFormula).
		Do()
	if len(records.Records) > 0 {
		fmt.Printf("已经存在")
		return status_code.ERROR_IS_USED
	}
	fmt.Printf("不存在")
	return status_code.SUCCESS
}

// 新增文件分类
func CreateCourse(data *Course) int {
	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"title": data.Title,
					"desc":  data.Desc,
				},
			},
		},
	}
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	_, err := table.AddRecords(recordsToSend)
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetCourseList() *airtable.Records {
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	records, err := table.GetRecords().ReturnFields("icon", "name", "title", "id").WithSort().Do()
	if err != nil {
		return nil
	}
	return records
}

// 编辑分类
func EditCourse(id string, data *JSONBindCourse) int {
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	record, err := table.GetRecord(id)

	var maps = make(map[string]interface{})
	if data.Title != nil {
		maps["title"] = data.Title
	}
	if data.Desc != nil {
		maps["desc"] = data.Desc
	}
	_, err = record.UpdateRecordPartial(maps)
	if err != nil {
		fmt.Printf(err.Error())
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 删除分类
func DeleteCourse(id string) int {
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	_, err := table.DeleteRecords([]string{id})
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

func GetCourseDetail(recordId string) (*airtable.Record, error) {
	table := airtableClient.GetTable(utils.AirtableDBId, "Course")
	return table.GetRecord(recordId)
}

//// 获取关联records
//func GetCourseLinkedRecords(recordId string) ([]*airtable.Record, []*airtable.Record, int) {
//	courseTable := airtableClient.GetTable(utils.AirtableDBId, "Course")
//	record, err := courseTable.GetRecord(recordId)
//	if err == nil {
//		fields := record.Fields
//		var posts []*airtable.Record
//		postTable := airtableClient.GetTable(utils.AirtableDBId, "Post")
//		for _, recordId := range fields["posts"].([]interface{}) {
//			linkedRecord, _ := postTable.GetRecord(recordId.(string))
//			posts = append(posts, linkedRecord)
//			fmt.Println(linkedRecord)
//		}
//
//		var columns []*airtable.Record
//		columnTable := airtableClient.GetTable(utils.AirtableDBId, "Column")
//		for _, recordId := range fields["columns"].([]interface{}) {
//			linkedRecord, _ := columnTable.GetRecord(recordId.(string))
//			columns = append(columns, linkedRecord)
//		}
//
//		return columns, posts, status_code.SUCCESS
//	} else {
//		return nil, nil, status_code.SUCCESS
//	}
//}
