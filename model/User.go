package model

import (
	"cant_forget/utils/status_code"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Username string     `gorm:"type:varchar(20);not null"`
	Password string     `gorm:"type:varchar(20);not null"`
	Email    *string    // A pointer to a string, allowing for null values
	Age      uint8      // An unsigned 8-bit integer
	Birthday *time.Time // A pointer to time.Time, can be null
	Avatar   string
	Role     int `gorm:"type:int"`
}

// 查询用户是否存在
func CheckUserExist(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return status_code.ERROR_USERNAME_USED // 1001
	}
	return status_code.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return status_code.ERROR
	}
	return status_code.SUCCESS
}

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
