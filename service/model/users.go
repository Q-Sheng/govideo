package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Salt     string `json:"salt" gorm:"column:salt"`
	Gender   int8   `json:"gender" gorm:"column:gender"`

	LiveInfo string `json:"liveInfo" gorm:"foreignKey:Uid"`
}
