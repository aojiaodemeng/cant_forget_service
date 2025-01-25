package model

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"fmt"
	"github.com/mehanizm/airtable"
	"gorm.io/gorm"
)

type Column struct {
	gorm.Model
	Id    string `json:"columnId" gorm:"not null;unique;primary_key;comment:专栏ID;size:90"`
	Title string `gorm:"type:varchar(20);not null"`
	Desc  string `gorm:"type:varchar(200)"`
	Img   string `gorm:"type:varchar(100)"`
}

// 新增
func CreateColumn(data *Column) int {
	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"title": data.Title,
					"desc":  data.Desc,
					"img":   data.Img,
				},
			},
		},
	}
	table := airtableClient.GetTable(utils.AirtableDBId, "Column")
	_, err := table.AddRecords(recordsToSend)
	if err != nil {
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetColumnList(courseUuid string) *airtable.Records {
	table := airtableClient.GetTable(utils.AirtableDBId, "Column")
	filterFormula := fmt.Sprintf("FIND('%s', {coursesStr}) > 0", courseUuid)
	records, err := table.GetRecords().WithFilterFormula(filterFormula).Do()
	if err != nil {
		print(err.Error())
		return nil
	}
	return records
}
