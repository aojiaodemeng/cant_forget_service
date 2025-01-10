package model

import (
	"cant_forget/utils/status_code"
	"gorm.io/gorm"
)

type PostType struct {
	gorm.Model
	Id   uint   `gorm:"not null;unique;primary_key;comment:帖子类型ID;size:90"`
	Name string `gorm:"type:varchar(20);not null"`
	Desc string `gorm:"type:varchar(200)"`
}

// 查询分类是否存在
func CheckPostTypeExist(name string) (code int) {
	var postType PostType
	db.Select("id").Where("name = ?", name).First(&postType)
	if postType.Id > 0 {
		return status_code.ERROR_POST_TYPE_USED
	}
	return status_code.SUCCESS
}

// 新增文件分类
func CreatePostType(data *PostType) int {
	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetFileCategories() []PostType {
	var fileCategories []PostType
	err = db.Find(&fileCategories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return fileCategories
}

// 编辑分类
func EditPostType(id int, data *PostType) int {
	var postType PostType
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&postType).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 删除分类
func DeletePostType(id string) int {
	var postType PostType
	err = db.Where("id=?", id).Delete(&postType).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}
