package model

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"encoding/base64"
	"fmt"
	"github.com/mehanizm/airtable"
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
	table := airtableClient.GetTable(utils.AirtableDBId, "User")
	filterFormula := fmt.Sprintf("AND({username}='%s')", name)
	records, _ := table.GetRecords().
		WithFilterFormula(filterFormula).
		Do()
	if len(records.Records) > 0 {
		fmt.Printf("已经存在")
		return status_code.ERROR_USERNAME_USED
	}
	fmt.Printf("不存在")
	return status_code.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"username": data.Username,
					"password": data.Password,
					"email":    data.Email,
					"age":      data.Age,
					"avatar":   data.Avatar,
				},
			},
		},
	}

	table := airtableClient.GetTable(utils.AirtableDBId, "User")
	_, err := table.AddRecords(recordsToSend)
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

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	code := CheckUserExist(username)

	if code == status_code.SUCCESS {
		return status_code.ERROR_USER_NOT_EXIST
	}

	if ScryptPw(password) != user.Password {
		return status_code.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return status_code.ERROR_USER_NO_RIGHT
	}
	return status_code.SUCCESS
}
