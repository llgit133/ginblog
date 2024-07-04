package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)" : json:"user_name"`
	Password string `gorm:"type:varchar(20)" : json:"password"`
	Role     int    `gorm:"type:int" : json:"role"`
}

// json 前后端数据交互
