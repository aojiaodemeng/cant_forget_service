package model

import (
	"cant_forget/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

// 连接配置数据库
func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	fmt.Printf(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, //表名使用单数模式，在复数模式下，结构体User对应的表名是users
	})
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数：", err)
	} else {
		fmt.Printf("连接数据库成功")
	}

	db.AutoMigrate(&User{}, &FileCategory{}, &CourseCategory{})

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)                  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) // SetConnMaxLifetime 设置了连接可复用的最大时间
}
