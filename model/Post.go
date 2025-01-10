package model

import (
	"cant_forget/utils/status_code"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID               uint             `gorm:"not null;unique;primary_key;comment:帖子ID;size:90"`
	Title            string           `gorm:"type:varchar(100);not null"`
	Content          string           `gorm:"type:longtext"`
	Img              string           `gorm:"type:varchar(100)"`
	Audio            string           `gorm:"type:varchar(100)"`
	Video            string           `gorm:"type:varchar(100)"`
	PostType         PostType         `gorm:"many2many:post_map;"`
	CourseCategories []CourseCategory `gorm:"many2many:post_map;"`
	Columns          []Column         `gorm:"many2many:post_map;"`
}

type CreatePostReqParams struct {
	Title       string `gorm:"type:varchar(20);not null"`
	Content     string `gorm:"type:longtext"`
	Img         string `gorm:"type:varchar(100)"`
	Audio       string `gorm:"type:varchar(100)"`
	Video       string `gorm:"type:varchar(100)"`
	PostTypeId  uint
	CategoryIds []uint
	ColumnIds   []uint
}

// 新增文章
func CreatePost(data *Post) int {

	// 创建一条文章记录，跳过所有关联字段
	// 备注：去掉.Omit(clause.Associations)的话，会在关联的表（比如Column表）里新建记录
	// 但是业务是，只是选择了（或者说关联了）标签、专栏、分类，而不是创建它们
	//err := db.Omit(clause.Associations).Create(&test).Error

	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR // 500
	}
	//// 添加关联
	////db.Model(&test).Association("Columns").Append(&Column{})
	//CreateMap(&data)

	return status_code.SUCCESS
}
