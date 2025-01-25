package model

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"fmt"
	"github.com/mehanizm/airtable"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID      string   `gorm:"not null;unique;primary_key;comment:帖子ID;size:90"`
	Title   string   `gorm:"type:varchar(100);not null"`
	Content string   `gorm:"type:longtext"`
	Img     string   `gorm:"type:varchar(100)"`
	Audio   string   `gorm:"type:varchar(100)"`
	Video   string   `gorm:"type:varchar(100)"`
	Courses []Course `gorm:"many2many:post_map;"`
	Columns []string `gorm:"many2many:post_map;"`
	Type    string
}

type CreatePostReqParams struct {
	Title   string `gorm:"type:varchar(20);not null"`
	Content string `gorm:"type:longtext"`
	Img     string `gorm:"type:varchar(100)"`
	//Audio       string `gorm:"type:varchar(100)"`
	//Video       string `gorm:"type:varchar(100)"`
	Type    string
	Columns []string
	Courses []string
	Efactor float32
}

// 新增文章
func CreatePost(data *CreatePostReqParams) int {
	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"title":   data.Title,
					"content": data.Content,
					"img":     data.Img,
					"Columns": data.Columns,
					"Courses": data.Courses,
					"type":    data.Type,
					"efactor": data.Efactor,
				},
			},
		},
	}

	table := airtableClient.GetTable(utils.AirtableDBId, "Post")
	_, err := table.AddRecords(recordsToSend)

	if err != nil {
		print(err.Error())
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询列表
func GetPostList(columnUuid string) *airtable.Records {
	table := airtableClient.GetTable(utils.AirtableDBId, "Post")
	filterFormula := fmt.Sprintf("FIND('%s', {columnsStr}) > 0", columnUuid)
	records, err := table.GetRecords().WithFilterFormula(filterFormula).Do()

	if err != nil {
		return nil
	}
	return records
}
