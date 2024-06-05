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
	Email    string `json:"email" gorm:"column:email"`

	Live_info string `json:"live_info" gorm:"column:live_info"`
}

type RegisterUser struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Salt     string `json:"salt" gorm:"column:salt"`
	//Gender   int8   `json:"gender" gorm:"column:gender"`
	Email   string `json:"email" binding:"required,email"`
	VerCode string `json:"verificationCode" binding:"required"`

	//Live_info string `json:"live_info" gorm:"column:live_info"`
}
