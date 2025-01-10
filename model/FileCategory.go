package model

import (
	"cant_forget/utils/status_code"
	"gorm.io/gorm"
)

type FileCategory struct {
	gorm.Model
	ID   string `gorm:"primary_key;"`
	Name string `gorm:"type:varchar(20);not null"`
}

// 查询分类是否存在
func CheckFileCategoryExist(id string) (code int) {
	var fileCategory FileCategory
	db.Select("id").Where("id = ?", id).First(&fileCategory)
	if fileCategory.ID != "" {
		return status_code.ERROR_FILE_CATEGORY_USED
	}
	return status_code.SUCCESS
}

// 新增文件分类
func CreateFileCategory(data *FileCategory) int {
	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetFileCategories() []FileCategory {
	var fileCategories []FileCategory
	err = db.Find(&fileCategories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return fileCategories
}

// 编辑分类
func EditFileCategory(id string, data *FileCategory) int {
	var fileCategory FileCategory
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&fileCategory).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 删除分类
func DeleteFileCategory(id string) int {
	var fileCategory FileCategory
	err = db.Where("id=?", id).Delete(&fileCategory).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}
