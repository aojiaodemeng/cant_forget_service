package model

import "gorm.io/gorm"

type PostMap struct {
	gorm.Model
	PostId           uint `gorm:"column:post_id"`
	CourseCategoryId uint `gorm:"column:course_category_id"`
	PostTypeId       uint `gorm:"column:post_type_id"`
	ColumnId         uint `gorm:"column:column_id"`
}
