package model

import (
	"cant_forget/utils/status_code"
	"gorm.io/gorm"
)

type Column struct {
	gorm.Model
	Id    uint   `json:"columnId" gorm:"not null;unique;primary_key;comment:专栏ID;size:90"`
	Title string `gorm:"type:varchar(20);not null"`
	Desc  string `gorm:"type:varchar(200)"`
	Img   string `gorm:"type:varchar(100)"`
}

// 新增
func CreateColumn(data *Column) int {
	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetColumnList() []PostType {
	var postTypes []PostType
	err = db.Find(&postTypes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return postTypes
}
