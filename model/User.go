package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string     `gorm:"type:varchar(20);not null" json:"username"`
	Password string     `gorm:"type:varchar(20);not null" json:"password"`
	Email    *string    // A pointer to a string, allowing for null values
	Age      uint8      // An unsigned 8-bit integer
	Birthday *time.Time // A pointer to time.Time, can be null
	Avatar   string
	Role     int `gorm:"type:int" json:"role"`
}
