package model

import (
	"cant_forget/utils/status_code"
	"gorm.io/gorm"
)

type CourseCategory struct {
	gorm.Model
	ID   string `gorm:"primary_key;"`
	Name string `gorm:"type:varchar(20);not null"`
}

// 查询分类是否存在
func CheckCourseCategoryExist(id string) (code int) {
	var courseCategory CourseCategory
	db.Select("id").Where("id = ?", id).First(&courseCategory)
	if courseCategory.ID != "" {
		return status_code.ERROR_FILE_CATEGORY_USED
	}
	return status_code.SUCCESS
}

// 新增文件分类
func CreateCourseCategory(data *CourseCategory) int {
	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询分类列表
func GetCourseCategories() []CourseCategory {
	var courseCategories []CourseCategory
	err = db.Find(&courseCategories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return courseCategories
}

// 编辑分类
func EditCourseCategory(id string, data *CourseCategory) int {
	var courseCategory CourseCategory
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&courseCategory).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 删除分类
func DeleteCourseCategory(id string) int {
	var courseCategory CourseCategory
	err = db.Where("id=?", id).Delete(&courseCategory).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}
